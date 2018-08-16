package xml

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

type Request struct {
	CountryCode string `json:"countryCode"`
}

type Response struct {
	XMLName     xml.Name `xml:"soapenv:Envelope"`
	CountryCode string   `xml:"soapenv:Body>hs:GetHolidaysAvailable>hs:countryCode"`
	Namespace   string   `xml:"xmlns:soapenv,attr"`
	NamespaceHs string   `xml:"xmlns:hs,attr"`
}

func (r Request) ToXML() Response {
	return Response{
		Namespace:   "http://schemas.xmlsoap.org/soap/envelope/",
		NamespaceHs: "http://www.holidaywebservice.com/HolidayService_v2/",
		CountryCode: r.CountryCode,
	}
}

func Test_Convert_JSON_Input_JSON_Should_Be_XML(t *testing.T) {
	expectedXML, _ := ioutil.ReadFile("./request.xml")
	var request Request
	jsonData := []byte(`{"countryCode":"UnitedStates"}`)
	json.Unmarshal(jsonData, &request)

	requestXML := request.ToXML()
	actualXML, _ := xml.MarshalIndent(requestXML, "", "\t")
	if string(expectedXML) != string(actualXML) {
		t.Errorf("expected \n%s but it got \n%s", expectedXML, actualXML)
	}
}
