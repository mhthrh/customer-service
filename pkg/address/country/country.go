package country

import (
	"bytes"
	"encoding/csv"
	"errors"
	customeError "github.com/mhthrh/common-lib/errors"
	countryError "github.com/mhthrh/common-lib/errors/country"
	"github.com/mhthrh/common-lib/model/address/country"
	csvFile "github.com/mhthrh/common-lib/pkg/util/file/csv"
)

const (
	path = "customer-service/file/countries/"
	name = "countries.csv"
)

type Countries struct {
	Countries []country.Country
}

func LoadCountries() (*Countries, *customeError.XError) {
	f := csvFile.New(path, name)

	bts, e := f.Read()
	if e != nil {
		return nil, countryError.FileUnreachable(customeError.RunTimeError(e))
	}
	reader := csv.NewReader(bytes.NewReader(bts))

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, countryError.FileUnreachable(customeError.RunTimeError(e))
	}
	if len(rows) < 1 {
		return nil, countryError.FileEmpty(customeError.RunTimeError(errors.New("no data")))
	}
	c := make([]country.Country, len(rows))
	for i, row := range rows {
		c[i] = country.Country{
			ID:   row[0],
			Name: row[1],
			Code: row[2],
		}
	}

	return &Countries{
		Countries: c,
	}, nil
}

func (c *Countries) FilterByCode(codes ...string) (Countries, *customeError.XError) {
	entry := make([]country.Country, 0)
	for _, code := range codes {
		for _, cnty := range c.Countries {
			if cnty.Code == code {
				entry = append(entry, cnty)
			}
		}
	}

	return Countries{Countries: entry}, nil
}
