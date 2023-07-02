package zImport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/kokizzu/gotro/L"

	"street/model/mProperty/wcProperty"

	"github.com/kokizzu/gotro/D/Tt"
)

// A Place Response
type PlaceResponse struct {
	Candidates []Candidate `json:"candidates"`
	Status     string      `json:"status"`
}

// Location
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// View Port
type ViewPort struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

// Geometry model
type Geometry struct {
	Location Location `json:"location"`
	ViewPort ViewPort `json:"viewport"`
}

type Candidate struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
}

func buildFullLocationSearchUrl(apiKey string, googleApiUrl string, address string) string {
	fields := "formatted_address%2Cname%2Crating%2Copening_hours%2Cgeometry"
	addressStr := address
	inputType := "textquery"
	googleApiKey := apiKey
	placeApiUrl := googleApiUrl + "/json?fields=" + fields + "&input=" + addressStr + "&inputtype=" + inputType + "&key=" + googleApiKey

	return placeApiUrl
}

func retrieveLatLongFromAddress(adapter *Tt.Adapter, apiKey string, googleApiUrl string) {
	defer subTaskPrint(`retrieveLatLongFromAddress: retrieve lat/long`)()

	propertyMutator := wcProperty.NewPropertyMutator(adapter)

	properties := propertyMutator.FindAllProperties()

	stat := &ImporterStat{Total: len(properties), PrintEvery: 10}
	for _, p := range properties {
		stat.Print()

		if p.Address == "" {
			stat.Skip()
			continue
		}
		if p.FormattedAddress != "" {
			stat.Skip()
			//log.Println("Property has address and lat/long already")
			continue
		}

		fullUrl := buildFullLocationSearchUrl(apiKey, googleApiUrl, p.Address)
		locationResponse, err := http.Get(fullUrl)

		if L.IsError(err, `retrieveLatLongFromAddress: get location response`) {
			stat.Ok(false)
			continue
		}

		responseData, err := io.ReadAll(locationResponse.Body)
		if L.IsError(err, `retrieveLatLongFromAddress: read response body`) {
			stat.Ok(false)
			continue
		}
		//fmt.Println("Response data => " + string(responseData))
		propertyLocation := PlaceResponse{}

		err = json.Unmarshal(responseData, &propertyLocation)
		if L.IsError(err, `retrieveLatLongFromAddress: unmarshal response data`) {
			stat.Ok(false)
			continue
		}
		if len(propertyLocation.Candidates) == 0 {
			stat.Skip()
			stat.Warn(`empty location`)
			//fmt.Println("There is no available location for this address")
			break
		}

		dataMutator := wcProperty.NewPropertyMutator(adapter)
		dataMutator.Property = *p
		dataMutator.Adapter = adapter
		dataMutator.FormattedAddress = propertyLocation.Candidates[0].FormattedAddress

		latitude := propertyLocation.Candidates[0].Geometry.Location.Lat
		longitude := propertyLocation.Candidates[0].Geometry.Location.Lng

		dataMutator.Coord = []any{latitude, longitude}
		dataMutator.UpdatedAt = time.Now().Unix()

		stat.Ok(dataMutator.DoOverwriteById())
	}

	stat.Print()
}

func ImportHouseLocation(adapter *Tt.Adapter) {
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	if googleApiKey == "" {
		fmt.Println("Require google api key to execute this operation")
		return
	}

	googleApiUrl := os.Getenv("GOOGLE_API_URL")
	if googleApiUrl == "" {
		fmt.Println("Require google api url to execute this operation")
		return
	}

	start := time.Now()
	retrieveLatLongFromAddress(adapter, googleApiKey, googleApiUrl)
	L.TimeTrack(start, "ImportHouseLocation")
}
