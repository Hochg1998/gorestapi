package model

//Product {id, category, image, name, price, isSale, createdAt, modifiedAt}
type Product struct {
	Id         int64    `json:"id"`
	CategoryId int64    `json:"categoryid"`
	Image      []string `json:"image"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	IsSale     bool     `json:"issale"`
	Rating     float32  `json:"rating"`
	CreatedAt  int      `json:"createdat"`
	ModifiedAt int      `json:"modifiedat"`
}
