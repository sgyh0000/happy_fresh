package models

type Cart struct {
	Id       Pk
	UserName string
	ItemId   Pk
	Num      uint64
}

type CartDetail struct {
	Id         Pk
	Item       Item
	Num        uint64
	TotalPrice Yuan
}

func FindCountByUser(username string) int64 {
	i, err := o.QueryTable(new(Cart)).Filter("user_name", username).Count()
	if err != nil {
		log.Error("error: ", err)
		return 0
	}
	return i
}

func FindByUser(username string) []CartDetail {
	var carts []CartDetail
	var cart []Cart
	_, err := o.QueryTable(new(Cart)).Filter("user_name", username).All(&cart)
	if err != nil {
		log.Error("error: ", err)
		return nil
	}
	for _, c := range cart {
		var item Item
		_ = o.QueryTable(new(Item)).Filter("id", c.ItemId).One(&item)
		carts = append(carts, CartDetail{
			Id:         c.Id,
			Item:       item,
			Num:        c.Num,
			TotalPrice: Yuan(c.Num * uint64(item.Price)),
		})
	}
	return carts
}

func DeleteByCartIds(ids []Pk) {
	_, _ = o.QueryTable(new(Cart)).Filter("id__in", ids).Delete()
}

func DeleteByCartId(id Pk) {
	_, _ = o.QueryTable(new(Cart)).Filter("id", id).Delete()
}

func AddItem(id Pk) uint64 {
	var cart Cart
	_ = o.QueryTable(new(Cart)).Filter("id", id).One(&cart)
	cart.Num++
	_, _ = o.Update(&cart)
	return cart.Num
}

func MinusItem(id Pk) uint64 {
	var cart Cart
	_ = o.QueryTable(new(Cart)).Filter("id", id).One(&cart)
	cart.Num--
	_, _ = o.Update(&cart)
	return cart.Num
}
