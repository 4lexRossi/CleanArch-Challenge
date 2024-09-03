package usecase

import (
	"github.com/4lexRossi/CleanArch-Challenge/internal/entity"
	"github.com/4lexRossi/CleanArch-Challenge/pkg/events"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var ordersResponse []OrderOutputDTO

	for _, order := range orders {
		orderResponse := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			FinalPrice: order.FinalPrice,
		}

		ordersResponse = append(ordersResponse, orderResponse)
	}

	return ordersResponse, nil
}
