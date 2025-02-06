package service

import (
	"github.com/mhthrh/GoNest/model/address"
	"github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
)

type Service struct {
}

func NewService() customer.ICustomer {
	return &Service{}
}
func (s Service) RegisterCustomer(address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetCustomerById(id string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetCustomerByName(name string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetCustomerByEmail(email string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetCustomerByPhone(phone string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ChangeStatus(id string, status customer.Status) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (s Service) EditCustomer(address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}
