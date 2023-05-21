package saAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	"database/sql"
	"time"

	"street/model/mAuth"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	chBuffer "github.com/kokizzu/ch-timed-buffer"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saAuth__ORM.GEN.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type saAuth__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type saAuth__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type saAuth__ORM.GEN.go
// go:generate msgp -tests=false -file saAuth__ORM.GEN.go -o saAuth__MSG.GEN.go

var userLogsDummy = UserLogs{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mAuth.TableUserLogs: func(tx *sql.Tx) *sql.Stmt {
		query := userLogsDummy.sqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}

type UserLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	RequestId  string      `json:"requestId,string" form:"requestId" query:"requestId" long:"requestId" msg:"requestId"`
	Error      string      `json:"error" form:"error" query:"error" long:"error" msg:"error"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	IpAddr4    string      `json:"ipAddr4" form:"ipAddr4" query:"ipAddr4" long:"ipAddr4" msg:"ipAddr4"`
	IpAddr6    string      `json:"ipAddr6" form:"ipAddr6" query:"ipAddr6" long:"ipAddr6" msg:"ipAddr6"`
	UserAgent  string      `json:"userAgent" form:"userAgent" query:"userAgent" long:"userAgent" msg:"userAgent"`
	Action     string      `json:"action" form:"action" query:"action" long:"action" msg:"action"`
	Traces     string      `json:"traces" form:"traces" query:"traces" long:"traces" msg:"traces"`
	StatusCode int16       `json:"statusCode" form:"statusCode" query:"statusCode" long:"statusCode" msg:"statusCode"`
}

func NewUserLogs(adapter *Ch.Adapter) *UserLogs {
	return &UserLogs{Adapter: adapter}
}

func (u UserLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mAuth.TableUserLogs
}

func (u *UserLogs) sqlTableName() string { //nolint:dupl false positive
	return `"userLogs"`
}

// insert, error if exists
func (u *UserLogs) sqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + u.sqlTableName() + `(` + u.sqlAllFields() + `) VALUES (?,?,?,?,?,?,?,?,?,?)`
}

func (u *UserLogs) sqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + u.sqlTableName()
}

func (u *UserLogs) sqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, requestId
	, error
	, actorId
	, ipAddr4
	, ipAddr6
	, userAgent
	, action
	, traces
	, statusCode
	`
}

func (u *UserLogs) sqlAllFields() string { //nolint:dupl false positive
	return `createdAt, requestId, error, actorId, ipAddr4, ipAddr6, userAgent, action, traces, statusCode`
}

func (u UserLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		u.CreatedAt,  // 0
		u.RequestId,  // 1
		u.Error,      // 2
		u.ActorId,    // 3
		u.IpAddr4,    // 4
		u.IpAddr6,    // 5
		u.UserAgent,  // 6
		u.Action,     // 7
		u.Traces,     // 8
		u.StatusCode, // 9
	}
}

func (u *UserLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (u *UserLogs) sqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (u *UserLogs) IdxRequestId() int { //nolint:dupl false positive
	return 1
}

func (u *UserLogs) sqlRequestId() string { //nolint:dupl false positive
	return `requestId`
}

func (u *UserLogs) IdxError() int { //nolint:dupl false positive
	return 2
}

func (u *UserLogs) sqlError() string { //nolint:dupl false positive
	return `error`
}

func (u *UserLogs) IdxActorId() int { //nolint:dupl false positive
	return 3
}

func (u *UserLogs) sqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (u *UserLogs) IdxIpAddr4() int { //nolint:dupl false positive
	return 4
}

func (u *UserLogs) sqlIpAddr4() string { //nolint:dupl false positive
	return `ipAddr4`
}

func (u *UserLogs) IdxIpAddr6() int { //nolint:dupl false positive
	return 5
}

func (u *UserLogs) sqlIpAddr6() string { //nolint:dupl false positive
	return `ipAddr6`
}

func (u *UserLogs) IdxUserAgent() int { //nolint:dupl false positive
	return 6
}

func (u *UserLogs) sqlUserAgent() string { //nolint:dupl false positive
	return `userAgent`
}

func (u *UserLogs) IdxAction() int { //nolint:dupl false positive
	return 7
}

func (u *UserLogs) sqlAction() string { //nolint:dupl false positive
	return `action`
}

func (u *UserLogs) IdxTraces() int { //nolint:dupl false positive
	return 8
}

func (u *UserLogs) sqlTraces() string { //nolint:dupl false positive
	return `traces`
}

func (u *UserLogs) IdxStatusCode() int { //nolint:dupl false positive
	return 9
}

func (u *UserLogs) sqlStatusCode() string { //nolint:dupl false positive
	return `statusCode`
}

func (u *UserLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		u.CreatedAt,  // 0
		u.RequestId,  // 1
		u.Error,      // 2
		u.ActorId,    // 3
		u.IpAddr4,    // 4
		u.IpAddr6,    // 5
		u.UserAgent,  // 6
		u.Action,     // 7
		u.Traces,     // 8
		u.StatusCode, // 9
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go
