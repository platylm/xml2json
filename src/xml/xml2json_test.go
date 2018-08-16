package xml

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func Test_Convert_XML_to_JSON_Should_Be_JSON(t *testing.T) {
	var actual GetCountriesAvailable
	xmlFile, _ := ioutil.ReadFile("./response.xml")
	expected := GetCountriesAvailable{
		CountryCode: []CountryCode{
			CountryCode{"Canada", "Canada"},
			CountryCode{"GreatBritain", "Great Britain and Wales"},
			CountryCode{"IrelandNorthern", "Northern Ireland"},
			CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
			CountryCode{"Scotland", "Scotland"},
			CountryCode{"UnitedStates", "United States"},
		},
	}
	xml.Unmarshal(xmlFile, &actual)
	for index, actualCountryCode := range actual.CountryCode {
		expectedCountryCode := expected.CountryCode[index]
		if expectedCountryCode != actualCountryCode {
			t.Errorf("expected at index: %d %s but it got %s", index, expectedCountryCode, actualCountryCode)
		}
	}

}

func Test_to_JSON_Should_Be_JSON(t *testing.T) {
	expected := `{"countrycode":[{"code":"Canada","description":"Canada"},{"code":"GreatBritain","description":"Great Britain and Wales"},{"code":"IrelandNorthern","description":"Northern Ireland"},{"code":"IrelandRepublicOf","description":"Republic of Ireland"},{"code":"Scotland","description":"Scotland"},{"code":"UnitedStates","description":"United States"}]}`
	getCountriesAvailable := GetCountriesAvailable{
		CountryCode: []CountryCode{
			CountryCode{"Canada", "Canada"},
			CountryCode{"GreatBritain", "Great Britain and Wales"},
			CountryCode{"IrelandNorthern", "Northern Ireland"},
			CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
			CountryCode{"Scotland", "Scotland"},
			CountryCode{"UnitedStates", "United States"},
		},
	}
	countriesRespone := getCountriesAvailable.ToJSON()
	actual, _ := json.Marshal(countriesRespone)
	if expected != string(actual) {
		t.Errorf("expected \n%s but got it \n%s", expected, actual)
	}
}
