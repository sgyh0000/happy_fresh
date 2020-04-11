package models

type Delivery struct {
	Id        Pk
	UserName  string // 关联的用户名
	RealName  string // 收件人姓名
	Address   string // 收货地址
	Code      string // 邮编
	Telephone string // 收货电话
}

func FindByUserName(username string) Delivery {

	var result Delivery
	err := o.QueryTable(new(Delivery)).Filter("user_name", username).One(&result)
	if err != nil {
		log.Error("error: ", err)
	}
	return result
}

func InsertOrUpdateDelivery(delivery Delivery) {
	existDelivery := FindByUserName(delivery.UserName)
	if existDelivery.Id > 0 {
		delivery.Id = existDelivery.Id
		_, _ = o.Update(&delivery)
	} else {
		_, _ = o.Insert(&delivery)
	}

}
