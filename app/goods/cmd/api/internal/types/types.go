// Code generated by goctl. DO NOT EDIT.
package types

type Goods struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListReq struct {
	PageNum  int64 `path:"pageNum,default=0"`
	PageSize int64 `path:"pageSize,default=10"`
}

type ListResp struct {
	Goods []Goods `json:"goods"`
}

type GoodsReq struct {
	Id int64 `path:"id"`
}

type GoodsResp struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Stocks int64   `json:"stocks"`
}
