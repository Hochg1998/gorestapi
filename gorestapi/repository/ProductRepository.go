package repository

import (
	"errors"
	"fmt"

	"github.com/ocg.com/go/01/lab/gorestapi/model"
)

type ProductRepo struct {
	products map[int64]*model.Product
	autoID   int64 //đây là biến đếm tự tăng gán giá trị cho id của product
}

var Products ProductRepo

func init() {
	Products = ProductRepo{autoID: 0}
	Products.products = make(map[int64]*model.Product)
	Products.InitData("sql:45312")
}

func (r *ProductRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ProductRepo) CreateNewProduct(product *model.Product) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	product.Id = nextID
	r.products[nextID] = product //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *ProductRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewProduct(&model.Product{
		CategoryId: 2,
		Image:      []string{"image 1", "image2"},
		Name:       "Herschel supply co 25l",
		Price:      75,
		IsSale:     true,
		Rating:     5.0,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000})

	r.CreateNewProduct(&model.Product{
		CategoryId: 1,
		Image:      []string{"image 3", "image 4"},
		Name:       "Denim jacket blue",
		Price:      92.5,
		IsSale:     false,
		Rating:     5.0,
		CreatedAt:  1610281342000,
		ModifiedAt: 1619283693000})
}

func (r *ProductRepo) GetAllProducts() map[int64]*model.Product {

	return r.products
}

func (r *ProductRepo) FindProductById(Id int64) (*model.Product, error) {
	if product, ok := r.products[Id]; ok {
		return product, nil //tìm được
	} else {
		return nil, errors.New("product not found")
	}
}

func (r *ProductRepo) DeleteProductById(Id int64) error {
	if _, ok := r.products[Id]; ok {
		delete(r.products, Id)
		return nil
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepo) UpdateProduct(product *model.Product) error {
	if _, ok := r.products[product.Id]; ok {
		r.products[product.Id] = product
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepo) Upsert(product *model.Product) int64 {
	if _, ok := r.products[product.Id]; ok {
		r.products[product.Id] = product //tìm thấy thì update
		return product.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewProduct(product)
	}
}

func (r *ProductRepo) Update(productId int64, averageRating float32) error {
	//TODO: cập nhật dữ liệu ở đây
	if _, ok := r.products[productId]; ok {
		r.products[productId].Rating = averageRating
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}
