package address

import (
	customerCity "customer-service/pkg/address/city"
	customerCountry "customer-service/pkg/address/country"
	"github.com/mhthrh/common-lib/model/address"
	"strings"
)

type Inter struct {
}
type Address struct {
	Adrs address.Address
}

func NewAddress(street, postalCode, state, cntry, cty string) (*Address, error) {
	if strings.Trim(street, " ") == "" {
		return nil, nil
	}
	if strings.Trim(postalCode, " ") == "" {
		return nil, nil
	}
	if strings.Trim(state, " ") == "" {
		return nil, nil
	}
	if strings.Trim(cntry, " ") == "" {

	}
	if strings.Trim(cty, " ") == "" {

	}
	c, e := customerCountry.LoadCountries()
	if e != nil {
		return nil, e
	}
	cResult, err := c.FilterByCode(cntry)
	if err != nil {
		return nil, err
	}
	if len(cResult.Countries) != 1 {
		return nil, nil
	}
	cy, e := customerCity.Load()
	if e != nil {
		return nil, e
	}
	cyResult := cy.FilterByCountry(cntry)

	if len(cyResult.Cities) != 1 {
		return nil, nil
	}

	return &Address{Adrs: address.Address{
		Street:     strings.Trim(street, " "),
		City:       cyResult.Cities[0],
		State:      state,
		PostalCode: postalCode,
		Country:    cResult.Countries[0],
	}}, nil
}
