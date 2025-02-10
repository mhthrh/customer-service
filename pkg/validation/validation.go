package validation

import (
	"context"
	"github.com/mhthrh/GoNest/model/address"
	"github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
)

type Validation struct {
}

func (v Validation) RegisterCustomer(ctx context.Context, address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (v Validation) GetCustomerById(ctx context.Context, id string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (v Validation) GetCustomerByName(ctx context.Context, name string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (v Validation) GetCustomerByEmail(ctx context.Context, email string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (v Validation) GetCustomerByPhone(ctx context.Context, phone string) (*customer.Customer, bool) {
	//TODO implement me
	panic("implement me")
}

func (v Validation) ChangeStatus(ctx context.Context, id string, status customer.Status) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (v Validation) EditCustomer(ctx context.Context, address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}
