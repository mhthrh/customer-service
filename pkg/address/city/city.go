package city

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/google/uuid"
	customeError "github.com/mhthrh/common-lib/errors"
	cityError "github.com/mhthrh/common-lib/errors/city"
	"github.com/mhthrh/common-lib/model/address/city"
	csvFile "github.com/mhthrh/common-lib/pkg/util/file/csv"
)

const (
	path = "customer-service/file/cities/"
	name = "cities.csv"
)

type City struct {
	Cities []city.City
}

func Load() (*City, *customeError.XError) {
	f := csvFile.New(path, name)
	bts, e := f.Read()
	if e != nil {
		return nil, cityError.FileUnreachable(customeError.RunTimeError(e))
	}

	reader := csv.NewReader(bytes.NewReader(bts))

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, cityError.FileUnreachable(customeError.RunTimeError(err))
	}
	c := make([]city.City, len(rows))
	for i, row := range rows {
		c[i] = city.City{
			ID:          uuid.New(),
			Name:        row[1],
			CountryCode: row[0],
		}
	}
	if len(c) == 0 {
		return nil, cityError.FileEmpty(customeError.RunTimeError(errors.New("no city found")))
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

func (c *City) FilterByCity(cti string) City {
	entry := make([]city.City, 0)

	for _, cnty := range c.Cities {
		if cnty.Name == cti {
			entry = append(entry, cnty)
		}
	}
	return City{Cities: entry}
}
func (c *City) FilterByCityAndCountry(cti, ctry string) City {
	entry := make([]city.City, 0)

	for _, cnty := range c.Cities {
		if cnty.Name == cti && cnty.CountryCode == ctry {
			entry = append(entry, cnty)
		}
	}
	return City{Cities: entry}
}
