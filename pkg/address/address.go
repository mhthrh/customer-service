package address

import (
	customerCity "customer-service/pkg/address/city"
	customerCountry "customer-service/pkg/address/country"
	"errors"
	customeError "github.com/mhthrh/common-lib/errors"
	addressError "github.com/mhthrh/common-lib/errors/address"
	"github.com/mhthrh/common-lib/model/address"
	"strings"
)

type Address struct {
	Adrs address.Address
}

func NewAddress(street, postalCode, state, cntry, cty string) (*Address, *customeError.XError) {
	if strings.Trim(street, " ") == "" {
		return nil, addressError.StreetNotFound(nil)
	}
	if strings.Trim(postalCode, " ") == "" {
		return nil, addressError.PostalCodeNotFound(nil)
	}
	if strings.Trim(state, " ") == "" {
		return nil, addressError.StateNotFound(nil)
	}
	if strings.Trim(cntry, " ") == "" {
		return nil, addressError.CountryNotFound(nil)
	}
	if strings.Trim(cty, " ") == "" {
		return nil, addressError.CityNotFound(nil)
	}
	c, e := customerCountry.LoadCountries()
	if e != nil {
		return nil, addressError.CountryNotFound(e)
	}
	cResult, err := c.FilterByCode(cntry)
	if err != nil {
		return nil, addressError.CityNotFound(err)
	}
	if len(cResult.Countries) != 1 {
		return nil, addressError.CountryNotFound(customeError.RunTimeError(errors.New("invalid countries length")))
	}
	cy, e := customerCity.Load()
	if e != nil {
		return nil, addressError.CityNotFound(e)
	}
	cyResult := cy.FilterByCityAndCountry(cty, cntry)

	if len(cyResult.Cities) != 1 {
		return nil, addressError.CityNotFound(customeError.RunTimeError(errors.New("count is more than 1")))
	}

	return &Address{Adrs: address.Address{
		Street:     strings.Trim(street, " "),
		City:       cyResult.Cities[0],
		State:      state,
		PostalCode: postalCode,
		Country:    cResult.Countries[0],
	}}, nil
}
