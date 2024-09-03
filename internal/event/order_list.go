package event

type OrderList struct {
	OrderList []interface{}
}

func (e *OrderList) GetResponse() []interface{} {
	return e.OrderList
}
