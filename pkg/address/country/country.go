package country

import (
	"bytes"
	"encoding/csv"
	"github.com/mhthrh/common-lib/model/address/country"
	csvFile "github.com/mhthrh/common-lib/pkg/util/file"
)

const (
	path = "customer-service/file/countries/"
	name = "countries.csv"
)

type Countries struct {
	Countries []country.Country
}

func LoadCountries() (*Countries, error) {
	f := csvFile.File{
		Name: name,
		Path: path,
		Data: nil,
	}
	e := f.Read()
	if e != nil {
		return nil, e
	}
	reader := csv.NewReader(bytes.NewReader(f.Data))

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
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

func (c *Countries) FilterByCode(codes ...string) (Countries, error) {
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
