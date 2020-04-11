package models

import (
	"time"
)

// 用户线索  请求的地址及请求时间
type Clue struct {
	Id         Pk
	Uri        string
	UserName   string
	DateCreate time.Time `orm:"auto_now_add;type(datetime)"`
}

func FindCluesByUserName(userName string) []Clue {

	var clues []Clue
	_, _ = o.QueryTable(new(Clue)).
		Filter("user_name", userName).
		Filter("uri__icontains", "/item/").
		Distinct().
		Limit(5, 0).All(&clues, "uri")
	return clues
}
