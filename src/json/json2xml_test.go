package json_test

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	. "json"
	"testing"
)

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
