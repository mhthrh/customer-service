package city

import (
	"bytes"
	"encoding/csv"
	"github.com/google/uuid"
	"github.com/mhthrh/common-lib/model/address/city"
	csvFile "github.com/mhthrh/common-lib/pkg/util/file"
)

const (
	path = "customer-service/file/cities/"
	name = "cities.csv"
)

type City struct {
	Cities []city.City
}

func Load() (*City, error) {
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
	c := make([]city.City, len(rows))
	for i, row := range rows {
		c[i] = city.City{
			ID:          uuid.New(),
			Name:        row[1],
			CountryCode: row[0],
		}
	}

	return &City{
		Cities: c,
	}, nil
}

func (c *City) FilterByCountry(country string) City {
	entry := make([]city.City, 0)

	for _, cnty := range c.Cities {
		if cnty.CountryCode == country {
			entry = append(entry, cnty)
		}
	}
	return City{Cities: entry}
}
