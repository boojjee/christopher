package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	// "log"
	"net/http"
)

type MyJsonName struct {
	AuthenticationResultCode string `json:"authenticationResultCode"`
	BrandLogoUri             string `json:"brandLogoUri"`
	Copyright                string `json:"copyright"`
	ResourceSets             []struct {
		EstimatedTotal float64 `json:"estimatedTotal"`
		Resources      []struct {
			Type    string `json:"__type"`
			Address struct {
				AdminDistrict    string `json:"adminDistrict"`
				CountryRegion    string `json:"countryRegion"`
				FormattedAddress string `json:"formattedAddress"`
			} `json:"address"`
			Bbox          []float64 `json:"bbox"`
			Confidence    string    `json:"confidence"`
			EntityType    string    `json:"entityType"`
			GeocodePoints []struct {
				CalculationMethod string    `json:"calculationMethod"`
				Coordinates       []float64 `json:"coordinates"`
				Type              string    `json:"type"`
				UsageTypes        []string  `json:"usageTypes"`
			} `json:"geocodePoints"`
			MatchCodes []string `json:"matchCodes"`
			Name       string   `json:"name"`
			Point      struct {
				Coordinates []float64 `json:"coordinates"`
				Type        string    `json:"type"`
			} `json:"point"`
		} `json:"resources"`
	} `json:"resourceSets"`
	StatusCode        float64 `json:"statusCode"`
	StatusDescription string  `json:"statusDescription"`
	TraceId           string  `json:"traceId"`
}

func GetProvinceFromBingMapByPoint(lat string, lon string) (string, string, error) {
	// request http api

	var (
		urlRequest    = `http://dev.virtualearth.net/REST/v1/Locations/`
		locationPoint = lat + `,` + lon
		fixParam      = `?includeEntityTypes=Neighborhood&o=json&key=`
		bingKey       = `AtxuaSKnnexxDeJ70Ha9ytdzzsvgB-793E03cm967jW5bPzg-Hj4nLUO-g3c5rdb`
		province      string
	)

	URI := urlRequest + locationPoint + fixParam + bingKey
	res, err := http.Get(URI)
	if err != nil {
		return "", "err", err
	}

	body, err := ioutil.ReadAll(res.Body)

	var jsbing MyJsonName
	json.Unmarshal(body, &jsbing)
	res.Body.Close()
	if err != nil {
		return "", "err", err
	}
	if res.StatusCode != 200 {
		return "", "err", errors.New("fail")
	}
	if jsbing.ResourceSets[0].EstimatedTotal == 0 {
		return "", "err", errors.New("fail")
	} else {
		province = jsbing.ResourceSets[0].Resources[0].Address.AdminDistrict
		return province, "", err
	}

}
