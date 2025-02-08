package dao

import (
	"context"
	"github.com/mhthrh/GoNest/model/address"
	"github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
)

type Dao struct {
}

func (d Dao) RegisterCustomer(ctx context.Context, address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerById(ctx context.Context, id string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerByName(ctx context.Context, name string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerByEmail(ctx context.Context, email string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) GetCustomerByPhone(ctx context.Context, phone string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (d Dao) ChangeStatus(ctx context.Context, id string, status customer.Status) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (d Dao) EditCustomer(ctx context.Context, address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}
