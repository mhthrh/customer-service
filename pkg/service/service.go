package service

import (
	"context"
	"customer-service/pkg/dao"
	"github.com/mhthrh/GoNest/model/address"
	"github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
	l "github.com/mhthrh/GoNest/pkg/logger"
)

type Service struct {
	db  dao.Dao
	log l.Log
}

func NewService(transactionId string, db dao.Dao) customer.ICustomer {

	return &Service{
		db:  db,
		log: l.Initialize(transactionId),
	}
}

func (s Service) RegisterCustomer(ctx context.Context, address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	s.log.Start()
	panic("implement me")
}

func (s Service) GetCustomerById(ctx context.Context, id string) (*customer.Customer, *cError.XError) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetCustomerByName(ctx context.Context, name string) (*customer.Customer, *cError.XError) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetCustomerByEmail(ctx context.Context, email string) (*customer.Customer, *cError.XError) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetCustomerByPhone(ctx context.Context, phone string) (*customer.Customer, *cError.XError) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ChangeStatus(ctx context.Context, id string, status customer.Status) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func (s Service) EditCustomer(ctx context.Context, address address.Address, customer customer.Customer) *cError.XError {
	//TODO implement me
	panic("implement me")
}
