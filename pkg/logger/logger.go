package logger

import (
	"github.com/mhthrh/common-lib/model/address"
	"github.com/mhthrh/common-lib/model/customer"
	cError "github.com/mhthrh/common-lib/model/error"
)

type Logger struct {
}

func NewLogger() customer.ICustomer {
	return &Logger{}
}
func (l Logger) RegisterCustomer(address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (l Logger) GetCustomerById(id string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) GetCustomerByName(name string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) GetCustomerByEmail(email string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) GetCustomerByPhone(phone string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) ChangeStatus(id string, status customer.Status) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (l Logger) EditCustomer(address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}
