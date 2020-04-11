package models

type Category struct {
	Id       Pk
	Name     string
	Desc     string
	Priority uint8
	Pic      string
}

type CategoryList struct {
	Category *Category
	Items    []*Item
	HotItems []*Item
}

func SelectAll() []*CategoryList {

	var categories []*Category
	_, err := o.QueryTable(new(Category)).OrderBy("priority").All(&categories)
	var result []*CategoryList
	if err != nil {
		log.Error("category - SelectAll", err)
		return nil
	}
	for _, c := range categories {
		items, success := FindByCategoryId(c.Id)
		if !success {
			return nil
		}
		hotItems, success := FindHotByCategoryId(c.Id)
		if !success {
			return nil
		}
		result = append(result, &CategoryList{
			Category: c,
			Items:    items,
			HotItems: hotItems,
		})
	}
	return result
}
