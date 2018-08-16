package xml2json_test

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

type Envelope struct {
	xmlName xml.Name `xml:Envelope`
	Body    Body
}

type Body struct {
	xmlName                       xml.Name `xml:Body`
	GetCountriesAvailableResponse GetCountriesAvailableResponse
}

type GetCountriesAvailableResponse struct {
	xmlName                     xml.Name `xml:GetCountriesAvailableResponse`
	GetCountriesAvailableResult GetCountriesAvailableResult
}

type GetCountriesAvailableResult struct {
	xmlName     xml.Name      `xml:GetCountriesAvailableResult`
	CountryCode []CountryCode `xml:GetCountriesAvailableResult>CountryCode`
}

type CountryCode struct {
	Code        string `xml:"Code"`
	Description string `xml:Description`
}

func Test_Convert_XML_to_JSON_Should_Be_JSON(t *testing.T) {
	var actual Envelope
	xmlFile, _ := ioutil.ReadFile("./response.xml")
	expected := Envelope{
		Body: Body{
			GetCountriesAvailableResponse: GetCountriesAvailableResponse{
				GetCountriesAvailableResult: GetCountriesAvailableResult{
					CountryCode: []CountryCode{
						CountryCode{"Canada", "Canada"},
						CountryCode{"GreatBritain", "Great Britain and Wales"},
						CountryCode{"IrelandNorthern", "Northern Ireland"},
						CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
						CountryCode{"Scotland", "Scotland"},
						CountryCode{"UnitedStates", "United States"},
					},
				},
			},
		},
	}
	xml.Unmarshal(xmlFile, &actual)
	for index, actualCountryCode := range actual.Body.GetCountriesAvailableResponse.
		GetCountriesAvailableResult.
		CountryCode {
		expectedCountryCode := expected.Body.GetCountriesAvailableResponse.GetCountriesAvailableResult.CountryCode[index]
		if expectedCountryCode != actualCountryCode {
			t.Errorf("expected at index: %d %s but it got %s", index, expectedCountryCode, actualCountryCode)
		}
	}

}
