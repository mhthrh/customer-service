package city

import (
	"customer-service/pkg/address/city"
	"customer-service/test"
	"testing"
)

var (
	c *city.City
	e error
)

func init() {
	c, e = city.Load()
}
func TestLoadCity(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "load test",
			Input:    nil,
			OutPut:   47868,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		if len(c.Cities) != tst.OutPut.(int) {
			t.Errorf("TestLoadCity failed for test %s", tst.Name)
		}
	}
}

func TestFilterByCode(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "filter by code",
			Input:    "IR",
			OutPut:   491,
			HasError: false,
			Err:      nil,
		}, {
			Name:     "filter by wrong code",
			Input:    "IROO",
			OutPut:   0,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		r := c.FilterByCountry(tst.Input.(string))
		if len(r.Cities) != tst.OutPut.(int) {
			t.Errorf("TestFilterByCode failed for test %s", tst.Name)
		}

	}
}
