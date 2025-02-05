package city

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/google/uuid"
	customeError "github.com/mhthrh/common-lib/errors"
	cityError "github.com/mhthrh/common-lib/errors/city"
	countryError "github.com/mhthrh/common-lib/errors/country"
	cityModel "github.com/mhthrh/common-lib/model/address/city"
	csvFile "github.com/mhthrh/common-lib/pkg/util/file/csv"
)

const (
	path = "customer-service/file/cities/"
	name = "cities.csv"
)

var (
	cities []cityModel.City
)

type City struct {
	path string
	name string
}

func New() cityModel.ICity {
	return City{
		path: path,
		name: name,
	}
}
func (c City) Load() *customeError.XError {
	f := csvFile.New(c.path, c.name)
	bts, e := f.Read()
	if e != nil {
		return cityError.FileUnreachable(customeError.RunTimeError(e))
	}

	reader := csv.NewReader(bytes.NewReader(bts))

	rows, err := reader.ReadAll()
	if err != nil {
		return cityError.FileUnreachable(customeError.RunTimeError(err))
	}
	cities = make([]cityModel.City, len(rows))
	for i, row := range rows {
		cities[i] = cityModel.City{
			ID:          uuid.New(),
			Name:        row[1],
			CountryCode: row[0],
		}
	}
	if len(cities) == 0 {
		return cityError.FileEmpty(customeError.RunTimeError(errors.New("no city found")))
	}

	return nil
}
func (c City) Cities() ([]cityModel.City, *customeError.XError) {
	if len(cities) == 0 {
		return nil, cityError.NotLoaded(nil)
	}
	return cities, nil
}
func (c City) GetByCountry(country string) ([]cityModel.City, *customeError.XError) {
	entry := make([]cityModel.City, 0)
	if len(cities) == 0 {
		return nil, countryError.NotLoaded(nil)
	}
	for _, cnty := range cities {
		if cnty.CountryCode == country {
			entry = append(entry, cnty)
		}
	}
	return entry, nil
}

func (c City) GetByCity(city string) ([]cityModel.City, *customeError.XError) {
	entry := make([]cityModel.City, 0)
	if len(cities) == 0 {
		return nil, cityError.NotLoaded(nil)
	}
	for _, cnty := range cities {
		if cnty.Name == city {
			entry = append(entry, cnty)
		}
	}
	return entry, nil
}

func (c City) GetByCityAndCountry(city, country string) (cityModel.City, *customeError.XError) {
	var entry cityModel.City
	if len(cities) == 0 {
		return cityModel.City{}, countryError.NotLoaded(nil)
	}
	for _, cnty := range cities {
		if cnty.Name == city && cnty.CountryCode == country {
			return entry, nil
		}
	}
	return cityModel.City{}, cityError.NotFound(nil, city, country)
}
