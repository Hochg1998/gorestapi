package model

//Product {id, category, image, name, price, isSale, createdAt, modifiedAt}
type Cart struct {
	Id         int64   `json:"id"`
	UserId     int64   `json:"userid"`
	Product    Product `json:"product"`
	Quantity   int8    `json:"quantity"`
	CreatedAt  int     `json:"createdat"`
	ModifiedAt int     `json:"modifiedat"`
}
