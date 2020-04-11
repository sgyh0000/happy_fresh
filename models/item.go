package models

import (
	"time"
)

var ItemDao = new(Item)

type Item struct {
	Id           Pk
	Name         string
	Desc         string
	Detail       string
	CategoryId   Pk
	CategoryName string
	Price        Yuan `orm:"digits(10);decimals(2)"`
	Unit         string
	Pic          string
	DetailPic    string
	IsHot        bool
	DateCreate   time.Time `orm:"auto_now_add;type(datetime)"`
}

func FindItemByIds(ids []Pk) ([]Item, bool) {

	var items []Item
	_, err := o.QueryTable(new(Item)).Filter("id__in", &ids).All(&items)
	if err != nil {
		return items, false
	}
	return items, true
}

func FindByCategoryId(cId Pk) ([]*Item, bool) {
	var items []*Item
	// 查询前5个
	_, err := o.QueryTable(new(Item)).Filter("category_id", cId).Limit(5, 0).All(&items)
	if err != nil {
		log.Error("err", err)
		return items, false
	}
	return items, true
}

func FindPageByCategoryId(cId Pk, page int64, pageSize int64) (Page, []Item) {
	var items []Item
	total, _ := o.QueryTable(new(Item)).Filter("category_id", cId).Count()
	_, _ = o.QueryTable(new(Item)).Filter("category_id", cId).Limit(pageSize, (page-1)*pageSize).All(&items)
	return NewPage(page, pageSize, total), items
}

func FindHotByCategoryId(cId Pk) ([]*Item, bool) {
	var items []*Item
	_, err := o.QueryTable(new(Item)).Filter("category_id", cId).Filter("is_hot", true).All(&items)
	if err != nil {
		log.Error("err", err)
		return items, false
	}
	return items, true
}

func FindByKeyWord(itemName string, page int64, pageSize int64) (Page, []Item) {
	var data []Item
	total, _ := o.QueryTable(new(Item)).Filter("name__icontains", itemName).Count()
	_, _ = o.QueryTable(new(Item)).Filter("name__icontains", itemName).Limit(pageSize, (page-1)*pageSize).All(&data)
	return NewPage(page, pageSize, total), data
}

func FindNewItems() []Item {
	var newItems []Item
	_, _ = o.QueryTable(new(Item)).OrderBy("-date_create").Limit(3, 0).All(&newItems)
	return newItems
}
