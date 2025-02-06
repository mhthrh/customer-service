package dao

import (
	"github.com/mhthrh/GoNest/model/address"
	"github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
)

type Dao struct {
}

func NewDao() customer.ICustomer {
	return &Dao{}
}
func (d Dao) RegisterCustomer(address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerById(id string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerByName(name string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerByEmail(email string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerByPhone(phone string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) ChangeStatus(id string, status customer.Status) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (d Dao) EditCustomer(address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}
