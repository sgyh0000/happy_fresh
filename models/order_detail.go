package models

type OrderDetail struct {
	Id      Pk
	OrderId Pk
	ItemId  Pk
	Num     uint64
	Price   Yuan
}

func FindByOrderId(orderId Pk) []OrderDetail {
	var orderDetails []OrderDetail
	_, _ = o.QueryTable(new(OrderDetail)).Filter("order_id", orderId).All(&orderDetails)
	return orderDetails
}
