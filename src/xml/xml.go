package xml

import "encoding/xml"

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

type CountriesRespone struct {
	Country []Country `json:"countrycode"`
}

type Country struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type GetCountriesAvailable struct {
	CountryCode []CountryCode `xml:"Body>GetCountriesAvailableResponse>GetCountriesAvailableResult>CountryCode" json:"countrycode"`
}

type CountryCode struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

func (g GetCountriesAvailable) ToJSON() CountriesRespone {
	countries := make([]Country, len(g.CountryCode))
	for index := range g.CountryCode {
		countries[index] = Country{
			Code:        g.CountryCode[index].Code,
			Description: g.CountryCode[index].Description,
		}
	}
	return CountriesRespone{
		Country: countries,
	}
}
