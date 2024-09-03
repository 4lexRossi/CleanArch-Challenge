package service

import (
	"context"

	"github.com/4lexRossi/CleanArch-Challenge/internal/infra/grpc/pb"
	"github.com/4lexRossi/CleanArch-Challenge/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.Blank) (*pb.OrderList, error) {
	orders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var protoOrders []*pb.Order
	for _, o := range orders {
		protoOrders = append(protoOrders, &pb.Order{
			Id:         o.ID,
			Price:      float32(o.Price),
			FinalPrice: float32(o.FinalPrice),
		})
	}

	return &pb.OrderList{Orders: protoOrders}, nil
}
