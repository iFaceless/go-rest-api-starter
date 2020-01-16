package rest

import "reflect"

type PaginationSchema struct {
	Total   int  `json:"total"`
	Offset  int  `json:"offset"`
	Limit   int  `json:"limit"`
	IsFirst bool `json:"is_first"`
	IsEnd   bool `json:"is_end"`
}

type Page struct {
	Pagination PaginationSchema `json:"pagination"`
	Data       interface{}      `json:"data"`
}

func NewPage(c *Context, data interface{}, total int) *Page {
	rv := reflect.TypeOf(data)
	if rv.Kind() != reflect.Slice {
		panic("page data must be of slice type")
	}

	offset, limit := c.Offset(), c.Limit()
	return &Page{
		Pagination: PaginationSchema{
			Total:   total,
			Offset:  offset,
			Limit:   limit,
			IsFirst: offset == 0,
			IsEnd:   offset+limit >= total,
		},
		Data: data,
	}
}
