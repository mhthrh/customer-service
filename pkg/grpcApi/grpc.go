package grpcApi

import (
	"context"
	"github.com/mhthrh/GoNest/model/customer/grpc/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type Customer struct {
	customer.UnimplementedCustomerServiceServer
	m sync.Locker
}

func (c Customer) RegisterCustomer(ctx context.Context, request *customer.Request) (*customer.Response, error) {
	log.Println(request.Customer)
	return &customer.Response{Customer: &customer.Customer{
		CustomerId: "123444",
		IdType:     1,
		UserName:   "",
		Password:   "",
		Email:      "",
		Mobile:     "",
		FirstName:  "mohsen",
		MiddleName: "taheri",
		LastName:   "",
		CreatedAt:  nil,
		UpdatedAt:  nil,
		Status:     0,
		Picture:    nil,
		Document:   nil,
	}}, status.Errorf(codes.OK, "")

}

func (c Customer) GetCustomerById(ctx context.Context, id *customer.Id) (*customer.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c Customer) GetCustomerByName(ctx context.Context, name *customer.Name) (*customer.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c Customer) mustEmbedUnimplementedCustomerServiceServer() {
	//TODO implement me
	panic("implement me")
}
