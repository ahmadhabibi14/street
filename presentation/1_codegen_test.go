package presentation

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
)

//go:generate go test -bench=BenchmarkGenerateViews

func BenchmarkGenerateViews(b *testing.B) {
	L.TIMETRACK_MIN_DURATION = 0

	p := codegen{
		ModelDir:  "../model",
		DomainDir: "../domain",
		SvelteDir: "../svelte",

		ProjectName: "street",

		ActionsGenFile:   "./actions.GEN.go",
		ApiRoutesGenFile: "./api_routes.GEN.go",
		JsApiGenFile:     "../svelte/jsApi.GEN.js",
		CmdRunGenFile:    "./cmd_run.GEN.go",
		WebViewGenFile:   "./web_view.GEN.go",
	}
	p.StartCodegen()

	b.SkipNow()
}

const NL = "\n"
const TAB = "\t"
const generatedComment = NL + `// Code generated by 1_codegen_test.go DO NOT EDIT.` + NL

type codegen struct {
	ModelDir string // ../model
	models   models

	DomainDir string // ../domain
	domains   domains

	SvelteDir string
	views     []string

	ActionsGenFile   string // actions.GEN.go
	ApiRoutesGenFile string // api_routes.GEN.go
	CmdRunGenFile    string // cmd_run.GEN.go
	JsApiGenFile     string // jsApi.GEN.js
	WebViewGenFile   string // web_view.GEN.go

	ProjectName string // based on go.mod
}

type models struct {
	types types

	lastPackageName string
}

func (m *models) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}
	if fileAst, ok := n.(*ast.File); ok {
		m.lastPackageName = fileAst.Name.Name
		return m
	}

	if typAst, ok := n.(*ast.TypeSpec); ok {
		name, fields := getStruct(typAst)
		m.types.AddType(m.lastPackageName+`.`+name, fields)
		return nil
	}
	return m
}

func (m *models) parseModelFile(path string) {
	out := L.ReadFile(path)
	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, path, out, parser.AllErrors)
	L.PanicIf(err, `parser.ParseFile: `+path)

	ast.Walk(m, p)
}

type tfield struct {
	Name string
	Type string
}

type tstruct struct {
	fields []tfield
}

func (s *tstruct) AddField(name, typ string) {
	s.fields = append(s.fields, tfield{
		Name: name,
		Type: typ,
	})
}

type types struct {
	byName map[string]tstruct
}

func (t *types) AddType(name string, s tstruct) {
	if t.byName == nil {
		t.byName = make(map[string]tstruct)
	}
	t.byName[name] = s
}

type consts struct {
	byName map[string]string
}

func (c *consts) AddConst(name, val string) {
	if c.byName == nil {
		c.byName = make(map[string]string)
	}
	c.byName[name] = val
}

type tmethod struct {
	MethodName string
	In         string
	Out        string
	Calls      []string
	Receiver   string
}

func (t *tmethod) Visit(node ast.Node) (w ast.Visitor) {
	// *ast.IfStmt
	// *ast.BlockStmt
	// *ast.AssignStmt
	// *ast.ExprStmt
	if callAst, ok := node.(*ast.CallExpr); ok {
		if funAst, ok := callAst.Fun.(*ast.SelectorExpr); ok {
			if identAst, ok := funAst.X.(*ast.Ident); ok {
				call := identAst.Name + `.` + funAst.Sel.Name
				t.Calls = append(t.Calls, call)
			}
		}
	}
	return t
}

func (t *tmethod) FullName() string {
	return t.Receiver + `.` + t.MethodName
}

type handlers struct {
	byRcvDotMethod map[string]tmethod
}

func (m *handlers) AddMethod(handler tmethod) {
	if m.byRcvDotMethod == nil {
		m.byRcvDotMethod = make(map[string]tmethod)
	}
	m.byRcvDotMethod[handler.FullName()] = handler
}

type domains struct {
	types       types
	consts      consts
	packageName string
	handlers    handlers
}

func (d *domains) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}
	if declAst, ok := n.(*ast.GenDecl); ok {
		switch declAst.Tok {
		// type ( blaIn ..., blaOut ... )
		case token.TYPE:
			for _, specAst := range declAst.Specs {
				// blaIn or blaOut declaration
				typAst := specAst.(*ast.TypeSpec)
				d.types.AddType(getStruct(typAst))
			}
			return nil
		case token.CONST:
			for _, specAst := range declAst.Specs {
				valAst := specAst.(*ast.ValueSpec)
				for z, constName := range valAst.Names {
					constVal := valAst.Values[z].(*ast.BasicLit).Value
					d.consts.AddConst(constName.Name, S.TrimChars(constVal, `"`+"`"))
				}
			}
			return nil
		case token.IMPORT:
			return nil
		}
	}
	// package
	if fileAst, ok := n.(*ast.File); ok {
		d.packageName = fileAst.Name.Name
		return d
	}
	// func or method declaration (handler)
	if funcAst, ok := n.(*ast.FuncDecl); ok {
		handler := getDomainHandler(funcAst)
		if handler != nil {
			d.handlers.AddMethod(*handler)
		}
		return nil
	}
	return d
}

func getStruct(typAst *ast.TypeSpec) (name string, fields tstruct) {
	structAst := typAst.Type.(*ast.StructType)
	vStruct := tstruct{}
	for _, fieldAst := range structAst.Fields.List {
		for _, name := range fieldAst.Names {
			vStruct.AddField(name.Name, getType(fieldAst.Type))
		}
	}
	return typAst.Name.Name, vStruct
}

func getDomainHandler(funcAst *ast.FuncDecl) *tmethod {
	res := &tmethod{}
	recv := funcAst.Recv
	if recv != nil &&
		len(recv.List) == 1 &&
		recv.List[0].Type != nil {
		if starAst, ok := recv.List[0].Type.(*ast.StarExpr); ok {
			res.Receiver = starAst.X.(*ast.Ident).Name
		}
	}
	res.MethodName = funcAst.Name.Name
	// not a handler, handler must only receive struct ends with 'In'
	if funcAst.Type.Params == nil ||
		len(funcAst.Type.Params.List) != 1 ||
		funcAst.Type.Params.List[0].Type == nil {
		return nil
	}
	res.In = getType(funcAst.Type.Params.List[0].Type)

	// not a handler, handler must only return something ends with 'Out'
	if funcAst.Type.Results == nil ||
		len(funcAst.Type.Results.List) != 1 ||
		funcAst.Type.Results.List[0].Type == nil {
		return nil
	}
	res.Out = getType(funcAst.Type.Results.List[0].Type)

	// not a handler, handler must have Domain receiver
	if funcAst.Recv == nil ||
		len(funcAst.Recv.List) != 1 ||
		len(funcAst.Recv.List[0].Names) != 1 {
		return nil
	}
	ast.Walk(res, funcAst.Body)
	return res

}

func getType(expr ast.Expr) string {
	// identifier, eg. uint64
	if identAst, ok := expr.(*ast.Ident); ok {
		return identAst.Name
	}
	// selector, eg. context.Context
	if selAst, ok := expr.(*ast.SelectorExpr); ok {
		lhs := selAst.X.(*ast.Ident).Name
		rhs := selAst.Sel.Name
		return lhs + `.` + rhs
	}
	// map, eg. map[string]any
	if mapAst, ok := expr.(*ast.MapType); ok {
		return `map[` + getType(mapAst.Key) + `]` + getType(mapAst.Value)
	}
	// star, eg. Tt.Adapter
	if starAst, ok := expr.(*ast.StarExpr); ok {
		return getType(starAst.X)
	}
	// array, eg. []string
	if arrayAst, ok := expr.(*ast.ArrayType); ok {
		return `[]` + getType(arrayAst.Elt)
	}
	panic(fmt.Sprintf(`unhandled type: %T %#v`, expr, expr))
}

func (c *codegen) StartCodegen() {

	// parse model files
	start := time.Now()
	err := filepath.Walk(c.ModelDir, func(path string, info os.FileInfo, err error) error {
		if L.IsError(err, `filepath.Walk`) {
			return err
		}
		if S.EndsWith(path, `.go`) &&
			!S.EndsWith(path, `_test.go`) {
			c.models.parseModelFile(path)
			_ = c.models.Visit // ^ called by this
		}
		return nil
	})
	L.PanicIf(err, `filepath.Walk Model`)
	L.TimeTrack(start, `parsing model dir`)

	// parse domain files
	start = time.Now()
	err = filepath.Walk(c.DomainDir, func(path string, info os.FileInfo, err error) error {
		if L.IsError(err, `filepath.Walk`) {
			return err
		}
		if S.EndsWith(path, `.go`) &&
			!S.EndsWith(path, `_test.go`) {
			c.domains.parseDomainFile(path)
			_ = c.domains.Visit // ^ called by this
		}
		return nil
	})
	L.PanicIf(err, `filepath.Walk Domain`)
	L.TimeTrack(start, `parsing domain dir`)

	c.GenerateActionsFile()
	c.GenerateApiRoutesFile()
	c.GenerateJsApiFile()
	c.GenerateCmdRunFile()

	// parse svelte files
	start = time.Now()
	err = filepath.Walk(c.SvelteDir, func(path string, info os.FileInfo, err error) error {
		if L.IsError(err, `filepath.Walk`) {
			return err
		}
		if S.EndsWith(path, `.svelte`) &&
			!S.Contains(path, `/_`) {
			c.views = append(c.views, path)
		}
		return nil
	})
	L.TimeTrack(start, `parsing svelte dir`)

	c.GenerateWebRouteFile()

}

func (d *domains) parseDomainFile(path string) {
	out := L.ReadFile(path)
	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, path, out, parser.AllErrors)
	L.PanicIf(err, `parser.ParseFile: `+path)

	ast.Walk(d, p)
}

func (d *domains) eachSortedHandler(eachFunc func(name string, handler tmethod)) {
	// sort handlers
	handlers := d.handlers.byRcvDotMethod
	handlerNames := make([]string, 0, len(handlers))
	byName := map[string]tmethod{}
	for _, handler := range handlers {
		// only add domain method, not session or any other struct
		if S.Contains(handler.Receiver, `Domain`) {
			handlerNames = append(handlerNames, handler.MethodName)
			byName[handler.MethodName] = handler
		}
	}
	sort.Strings(handlerNames)

	for _, name := range handlerNames {
		eachFunc(name, byName[name])
	}
}

func (c *codegen) GenerateActionsFile() {
	defer L.TimeTrack(time.Now(), `GenerateActionsFile`)

	b := bytes.Buffer{}
	b.WriteString(`package presentation
` + generatedComment + `
import (
	"` + c.ProjectName + `/domain"
)

var allCommands = []string{
`)
	c.domains.eachSortedHandler(func(name string, handler tmethod) {
		b.WriteString(TAB + `domain.` + name + `Action,` + NL)
	})
	b.WriteString(`}` + NL)
	b.WriteString(generatedComment)

	L.CreateFile(c.ActionsGenFile, b.String())
}

func (c *codegen) GenerateApiRoutesFile() {
	defer L.TimeTrack(time.Now(), `GenerateApiRoutesFile`)

	b := bytes.Buffer{}
	b.WriteString(`package presentation
` + generatedComment + `
import (
	"context"

	"github.com/gofiber/fiber/v2"

	"street/domain"
)

func ApiRoutes(fw *fiber.App, d *domain.Domain) {
`)
	c.domains.eachSortedHandler(func(name string, handler tmethod) {
		b.WriteString(`
	// ` + name + `
	fw.Post("/"+domain.` + name + `Action, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.` + name + `In{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.` + name + `Action); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.` + name + `(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out)
	})
`)
	})

	b.WriteString(NL + `}` + NL)
	b.WriteString(generatedComment)

	L.CreateFile(c.ApiRoutesGenFile, b.String())
}

func (c *codegen) GenerateJsApiFile() {
	defer L.TimeTrack(time.Now(), `GenerateJsApiFile`)

	b := bytes.Buffer{}

	b.WriteString(`
import axios from "axios";` + NL)
	b.WriteString(generatedComment)

	c.domains.eachSortedHandler(func(name string, handler tmethod) {
		fields := c.domains.types.byName[handler.In].fields
		c.jsObject(&b, name+`In`, fields, 0)

		fields = c.domains.types.byName[handler.Out].fields
		c.jsObject(&b, name+`Out`, fields, 0)

		c.jsFunc(&b, handler)
	})

	b.WriteString(generatedComment)

	L.CreateFile(c.JsApiGenFile, b.String())
}

func (c *codegen) jsObject(b *bytes.Buffer, name string, fields []tfield, indent int) {
	b.WriteString(`/**
 * @typedef {Object} ` + name + NL)
	c.jsDoc(b, fields, ``)
	b.WriteString(` */
const ` + name + ` = {
`)
	c.jsFields(b, fields, indent)
	b.WriteString(`}
`)
}

func (c *codegen) jsFields(content *bytes.Buffer, fields []tfield, indent int) {
	for _, field := range fields {
		c.jsField(content, field, indent)
	}
}

func (c *codegen) jsField(b *bytes.Buffer, field tfield, indent int) {
	t := field.Type
	// skip unecessary fields
	switch t {
	case `Tt.Adapter`, `Ch.Adapter`:
		return
	}

	// write field names
	spaces := S.Repeat(` `, indent*2)
	b.WriteString(`  ` + spaces + field.Name + `: `)

	// write field type
	switch t {
	case `int`, `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64`, `float32`, `float64`:
		b.WriteString(`0,`)
	case `string`:
		b.WriteString(`'',`)
	default:
		ty := c.models.types.byName[t]
		b.WriteString(`{ // ` + t + NL)
		c.jsFields(b, ty.fields, indent+1)
		b.WriteString(`  ` + spaces + `},`)
	}
	b.WriteString(` // ` + t + NL)
}

func (c *codegen) jsDoc(b *bytes.Buffer, fields []tfield, parent string) {
	for _, field := range fields {
		t := field.Type
		jsT := t
		switch t {
		case `Tt.Adapter`, `Ch.Adapter`:
			continue
		case `int`, `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64`, `float32`, `float64`:
			jsT = `number`
		case `string`:
			jsT = `String`
		default:
			ty := c.models.types.byName[t]
			if len(ty.fields) > 0 {
				c.jsDoc(b, ty.fields, field.Name+`.`)
				continue
			}
			jsT = `Object`
		}
		b.WriteString(` * @property {` + jsT + `} ` + parent + field.Name + NL)
	}
}

func (c *codegen) jsFunc(b *bytes.Buffer, handler tmethod) {
	action := c.domains.consts.byName[handler.MethodName+`Action`]
	b.WriteString(`/**
 * @callback ` + handler.MethodName + `Callback
 * @param {` + handler.Out + `} o
 * @returns {Promise}
 */
/**
 * @param  {` + handler.In + `} i
 * @param {` + handler.MethodName + `Callback} cb
 * @returns {Promise}
 */
async function ` + handler.MethodName + `( i, cb ) {
  return await axios.post( '/` + action + `', i ).then( cb )
}` + NL)

	b.WriteString(NL)
}

func (c *codegen) GenerateCmdRunFile() {
	defer L.TimeTrack(time.Now(), `GenerateCmdRunFile`)

	b := bytes.Buffer{}
	b.WriteString(`package presentation

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"

	"street/domain"
)

` + generatedComment + `

func cmdRun(b *domain.Domain, action string, payload []byte) {
	switch action {`)
	c.domains.eachSortedHandler(func(name string, handler tmethod) {
		b.WriteString(`

	case domain.` + name + `Action:
		in := domain.` + name + `In{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.` + name + `(&in)
		fmt.Println(X.ToJsonPretty(out))
`)
	})
	b.WriteString(NL + TAB + `}` + NL + `}` + NL)
	b.WriteString(generatedComment)

	L.CreateFile(c.CmdRunGenFile, b.String())
}

func (c *codegen) GenerateWebRouteFile() {
	L.TimeTrack(time.Now(), `GenerateWebRouteFile`)

	b := bytes.Buffer{}
	b.WriteString(`package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)

` + generatedComment + `

var viewList = map[string]string{
`)
	sort.Strings(c.views)
	var cacheNames []string
	for _, svelteFile := range c.views {
		left := S.LeftOfLast(svelteFile, `.svelte`)
		htmlFile := left + `.html`

		cacheName := left[len(c.SvelteDir):]
		cacheName = S.Replace(cacheName, `/`, ` `)
		cacheName = S.PascalCase(cacheName)

		b.WriteString(TAB + S.BT(cacheName) + `: ` + S.BT(htmlFile) + `, // ` + svelteFile + NL)
		cacheNames = append(cacheNames, cacheName)
	}
	b.WriteString(`}` + NL + NL)

	for _, cacheName := range cacheNames {
		b.WriteString(`
func (v *Views) Render` + cacheName + `(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[` + S.BT(cacheName) + `].Str(m))
}` + NL)

	}

	b.WriteString(generatedComment)

	L.CreateFile(c.WebViewGenFile, b.String())
}
