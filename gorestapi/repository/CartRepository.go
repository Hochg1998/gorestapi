package repository

import (
	"errors"
	"fmt"

	"github.com/ocg.com/go/01/lab/gorestapi/model"
)

type CartRepo struct {
	carts  map[int64]*model.Cart
	autoID int64 //đây là biến đếm tự tăng gán giá trị cho id của Cart
}

var Carts CartRepo

func init() {
	Carts = CartRepo{autoID: 0}
	Carts.carts = make(map[int64]*model.Cart)
	Carts.InitData("sql:45312")
}

func (r *CartRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *CartRepo) CreateNewCart(cart *model.Cart) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	cart.Id = nextID
	r.carts[nextID] = cart //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *CartRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewCart(&model.Cart{
		Id:     1,
		UserId: 1,
		Product: model.Product{
			Id:         1,
			CategoryId: 2,
			Image:      []string{"image 1", "image2"},
			Name:       "Herschel supply co 25l",
			Price:      75,
			IsSale:     true,
			Rating:     5.0,
			CreatedAt:  1614362898000,
			ModifiedAt: 1615410795000},
		Quantity:   2,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})

	r.CreateNewCart(&model.Cart{
		Id:     2,
		UserId: 1,
		Product: model.Product{
			Id:         2,
			CategoryId: 1,
			Image:      []string{"image 3", "image 4"},
			Name:       "Denim jacket blue",
			Price:      92.5,
			Rating:     5.0,
			CreatedAt:  1610281342000,
			ModifiedAt: 1619283693000},
		Quantity:   1,
		CreatedAt:  1623353412432,
		ModifiedAt: 1623353412432,
	})
}

func (r *CartRepo) GetAllCarts() map[int64]*model.Cart {
	return r.carts
}

func (r *CartRepo) FindCartById(Id int64) (*model.Cart, error) {
	if cart, ok := r.carts[Id]; ok {
		return cart, nil //tìm được
	} else {
		return nil, errors.New("cart not found")
	}
}

func (r *CartRepo) DeleteCartById(Id int64) error {
	if _, ok := r.carts[Id]; ok {
		delete(r.carts, Id)
		return nil
	} else {
		return errors.New("cart not found")
	}
}

func (r *CartRepo) UpdateCart(cart *model.Cart) error {
	if _, ok := r.carts[cart.Id]; ok {
		r.carts[cart.Id] = cart
		return nil //tìm được
	} else {
		return errors.New("cart not found")
	}
}

func (r *CartRepo) Upsert(cart *model.Cart) int64 {
	if _, ok := r.carts[cart.Id]; ok {
		r.carts[cart.Id] = cart //tìm thấy thì update
		return cart.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewCart(cart)
	}
}
