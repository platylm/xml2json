package xml

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
