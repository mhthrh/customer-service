package address

import (
	customerCity "customer-service/pkg/address/city"
	customerCountry "customer-service/pkg/address/country"
	"errors"
	customeError "github.com/mhthrh/common-lib/errors"
	"github.com/mhthrh/common-lib/model/address"
	"strings"
)

type Inter struct {
}
type Address struct {
	Adrs address.Address
}

func NewAddress(street, postalCode, state, cntry, cty string) (*Address, *customeError.XError) {
	if strings.Trim(street, " ") == "" {
		return nil, customeError.StreetNotFound(nil)
	}
	if strings.Trim(postalCode, " ") == "" {
		return nil, customeError.PostalCodeNotFound(nil)
	}
	if strings.Trim(state, " ") == "" {
		return nil, customeError.StateNotFound(nil)
	}
	if strings.Trim(cntry, " ") == "" {
		return nil, customeError.CountryNotFound(nil)
	}
	if strings.Trim(cty, " ") == "" {
		return nil, customeError.CityNotFound(nil)
	}
	c, e := customerCountry.LoadCountries()
	if e != nil {
		return nil, customeError.CountryNotFound(e)
	}
	cResult, err := c.FilterByCode(cntry)
	if err != nil {
		return nil, customeError.CityNotFound(err)
	}
	if len(cResult.Countries) != 1 {
		return nil, customeError.CountryNotFound(errors.New("count is more than 1"))
	}
	cy, e := customerCity.Load()
	if e != nil {
		return nil, customeError.CityNotFound(e)
	}
	cyResult := cy.FilterByCountry(cntry)

	if len(cyResult.Cities) != 1 {
		return nil, customeError.CityNotFound(errors.New("count is more than 1"))
	}

	return &Address{Adrs: address.Address{
		Street:     strings.Trim(street, " "),
		City:       cyResult.Cities[0],
		State:      state,
		PostalCode: postalCode,
		Country:    cResult.Countries[0],
	}}, nil
}
