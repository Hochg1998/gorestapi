package repository

import (
	"errors"
	"fmt"

	"github.com/ocg.com/go/01/lab/gorestapi/model"
)

type PriceRepo struct {
	prices map[int64]*model.Price
	autoID int64 //đây là biến đếm tự tăng gán giá trị cho id của price
}

var Prices PriceRepo

func init() {
	Prices = PriceRepo{autoID: 0}
	Prices.prices = make(map[int64]*model.Price)
	Prices.InitData("sql:45312")
}

func (r *PriceRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *PriceRepo) CreateNewPrice(price *model.Price) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	price.Id = nextID
	r.prices[nextID] = price //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *PriceRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewPrice(&model.Price{
		Price:      75,
		ProductId:  1,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000})

	r.CreateNewPrice(&model.Price{
		Price:      92.5,
		ProductId:  2,
		CreatedAt:  1610281342000,
		ModifiedAt: 1619283693000})
}

func (r *PriceRepo) GetAllPrices() map[int64]*model.Price {
	return r.prices
}

func (r *PriceRepo) FindPriceById(Id int64) (*model.Price, error) {
	if price, ok := r.prices[Id]; ok {
		return price, nil //tìm được
	} else {
		return nil, errors.New("price not found")
	}
}

func (r *PriceRepo) DeletePriceById(Id int64) error {
	if _, ok := r.prices[Id]; ok {
		delete(r.prices, Id)
		return nil
	} else {
		return errors.New("price not found")
	}
}

func (r *PriceRepo) UpdatePrice(price *model.Price) error {
	if _, ok := r.prices[price.Id]; ok {
		r.prices[price.Id] = price
		return nil //tìm được
	} else {
		return errors.New("price not found")
	}
}

func (r *PriceRepo) Upsert(price *model.Price) int64 {
	if _, ok := r.prices[price.Id]; ok {
		r.prices[price.Id] = price //tìm thấy thì update
		return price.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewPrice(price)
	}
}
