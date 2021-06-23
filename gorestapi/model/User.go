package model

type User struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	UserName   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Birthday   string `json:"birthday"`
	Status     bool   `json:"status"`
	CreatedAt  int    `json:"createdat"`
	ModifiedAt int    `json:"modifiedat"`
}
