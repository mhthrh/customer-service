package test

import (
	"customer-service/pkg/address/country"
	"testing"
)

func TestLoadCountry(t *testing.T) {
	c, e := country.LoadCountries()
	if e != nil {
		t.Error(e)
	}
	if len(c.Countries) < 10 {
		t.Error("countries length should be greater than 10")
	}
	if c.Countries[0].Code != "AF" {
		t.Error("code should be AF")
	}
}
