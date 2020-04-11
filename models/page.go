package models

type Page struct {
	Page     int64
	PageSize int64
	Total    int64
	HasNext  bool
	HasPre   bool
}

func NewPage(page int64, pageSize int64, total int64) Page {
	return Page{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		HasNext:  total > page*pageSize,
		HasPre:   page > 1,
	}
}
