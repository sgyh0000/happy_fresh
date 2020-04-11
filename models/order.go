package models

import (
	"time"
)

type OrderStatus uint

const (
	_ = iota
	UnPay
	PAID
	CANCELED
	INVALID
)

var OrderStatusName = map[OrderStatus]string{
	UnPay:    "未支付",
	PAID:     "已支付",
	CANCELED: "已取消",
	INVALID:  "非法订单",
}

type Order struct {
	Id         Pk
	UserName   string
	OrderDate  time.Time `orm:"auto_now_add;type(datetime)"`
	Status     OrderStatus
	TotalPrice Yuan
}

type OrderVo struct {
	User         User
	Order        Order
	OrderStatus  string
	Items        []Item
	OrderDetails []OrderDetail
}

func AddOrder(cartId []Pk, userName string) {

	carts := FindByUser(userName)
	var totalPrice Yuan = 0
	for _, cart := range carts {
		totalPrice += cart.TotalPrice
	}
	var order = &Order{
		UserName:   userName,
		Status:     UnPay,
		TotalPrice: totalPrice,
	}
	orderId, _ := o.Insert(order)
	i, _ := o.QueryTable(new(OrderDetail)).PrepareInsert()
	for _, cart := range carts {
		orderDetail := &OrderDetail{
			OrderId: Pk(orderId),
			ItemId:  cart.Item.Id,
			Num:     cart.Num,
			Price:   cart.TotalPrice,
		}
		_, _ = i.Insert(orderDetail)
	}
}

func FindOrderByUserName(username string) []OrderVo {
	var orderVos []OrderVo
	var orders []Order
	_, _ = o.QueryTable(new(Order)).Filter("user_name", username).All(&orders)

	user, _ := FindUserByUserName(username)

	for _, order := range orders {
		var orderVo OrderVo
		orderVo.Order = order
		orderVo.User = user
		orderVo.OrderDetails = FindByOrderId(order.Id)
		orderVo.OrderStatus = OrderStatusName[order.Status]

		var itemIds []Pk

		for _, orderDetail := range orderVo.OrderDetails {
			itemIds = append(itemIds, orderDetail.ItemId)
		}

		orderVo.Items, _ = FindItemByIds(itemIds)

		orderVos = append(orderVos, orderVo)
	}

	return orderVos
}
