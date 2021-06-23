package model

type Review struct {
	Id        int64  `json:"Id"`
	ProductId int64  `json:"ProductId"`
	Comment   string `json:"comment"`
	Rating    int    `json:"rating"`
}
