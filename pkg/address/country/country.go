package country

import (
	"bytes"
	"encoding/csv"
	"errors"
	customeError "github.com/mhthrh/common-lib/errors"
	countryError "github.com/mhthrh/common-lib/errors/country"
	cModel "github.com/mhthrh/common-lib/model/address/country"
	csvFile "github.com/mhthrh/common-lib/pkg/util/file/csv"
)

const (
	path = "customer-service/file/countries/"
	name = "countries.csv"
)

var (
	countries []cModel.Country
)

type Country struct {
	path string
	name string
}

func New() cModel.ICountry {
	return &Country{
		path: path,
		name: name,
	}
}

func (c Country) Load() *customeError.XError {
	f := csvFile.New(c.path, c.name)

	bts, e := f.Read()
	if e != nil {
		return countryError.FileUnreachable(customeError.RunTimeError(e))
	}
	reader := csv.NewReader(bytes.NewReader(bts))

	rows, err := reader.ReadAll()
	if err != nil {
		return countryError.FileUnreachable(customeError.RunTimeError(e))
	}
	if len(rows) < 1 {
		return countryError.FileEmpty(customeError.RunTimeError(errors.New("no data")))
	}
	countries = make([]cModel.Country, len(rows))
	for i, row := range rows {
		countries[i] = cModel.Country{
			ID:   row[0],
			Name: row[1],
			Code: row[2],
		}
	}

	return nil
}

func (c Country) Countries() ([]cModel.Country, *customeError.XError) {
	if len(countries) == 0 {
		return nil, countryError.NotLoaded(nil)
	}
	return countries, nil
}

func (c Country) GetByName(name string) (*cModel.Country, *customeError.XError) {
	for _, cnty := range countries {
		if cnty.Name == name {
			return &cnty, nil
		}
	}

	return nil, countryError.NotFound(nil, name)
}

func (c Country) GetByCode(code string) (*cModel.Country, *customeError.XError) {
	for _, cnty := range countries {
		if cnty.Code == code {
			return &cnty, nil
		}
	}

	return nil, countryError.NotFound(nil, code)
}
