package model

type Price struct {
	Id         int64   `json:"id"`
	Price      float64 `json:"price"`
	ProductId  int64   `json:"productid"`
	CreatedAt  int     `json:"createdat"`
	ModifiedAt int     `json:"modifiedat"`
}
