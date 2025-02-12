package grpc

import (
	"context"
	"fmt"
	"github.com/mhthrh/GoNest/model/customer/github.com/mhthrh/GoNest/model/customer"
)

type Customer struct {
}

func (c Customer) RegisterCustomer(ctx context.Context, request *customer.Request) (*customer.Response, error) {
	address := request.GetCustomer()
	customer1 := request.GetCustomer()
	fmt.Println(address, customer1)
	return nil, nil
}

func (c Customer) mustEmbedUnimplementedServiceServer() {
	//TODO implement me
	panic("implement me")
}
