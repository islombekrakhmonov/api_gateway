package client

import (
	"api-gateway/config"
	"api-gateway/genproto/order_service"
	"api-gateway/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type ServiceManagerI interface {
	OrderService() order_service.OrderServiceServer
	UserService() user_service.UserServiceServer
	ProductService() order_service.ProductServiceClient
}

type grpcClient struct {
	orderService order_service.OrderServiceServer
	userService user_service.UserServiceServer
	productService order_service.ProductServiceServer
}

func NewGrpcClient(cfg config.Config) (ServiceManagerI, error){
	connOrderService , err := grpc.Dial(
		cfg.OrderServiceHost + cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err 
	}

	connUserService , err := grpc.Dial(cfg.UserServiceHost + cfg.UserServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connProductService, err := grpc.Dial(cfg.OrderServiceHost + cfg.OrderServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	} 
	return &grpcClient{
		orderService: order_service.NewOrderServiceClient(connOrderService),
		userService: user_service.NewUserServiceClient(connUserService),
		connProductService: order_service.NewProductServiceClient(connProductService),
	}
}


func (g *grpcClient) OrderService() order_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClient) UserService() user_service.UserServiceClient {
	return g.userService
}

func (g *grpcClient) ProductService() order_service.ProductServiceClient{
	return g.productService
}