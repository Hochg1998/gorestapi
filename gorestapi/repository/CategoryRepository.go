package repository

import (
	"errors"
	"fmt"

	"github.com/ocg.com/go/01/lab/gorestapi/model"
)

type CategoryRepo struct {
	categories map[int64]*model.Category
	autoID     int64 //đây là biến đếm tự tăng gán giá trị cho id của category
}

var Categories CategoryRepo

func init() {
	Categories = CategoryRepo{autoID: 0}
	Categories.categories = make(map[int64]*model.Category)
	Categories.InitData("sql:45312")
}

func (r *CategoryRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *CategoryRepo) CreateNewCategory(category *model.Category) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	category.Id = nextID
	r.categories[nextID] = category //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *CategoryRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewCategory(&model.Category{
		Name: "Women",
	})

	r.CreateNewCategory(&model.Category{
		Name: "Men",
	})

	r.CreateNewCategory(&model.Category{
		Name: "Kids",
	})

	r.CreateNewCategory(&model.Category{
		Name: "Boys",
	})

}

func (r *CategoryRepo) GetAllCategories() map[int64]*model.Category {
	return r.categories
}

func (r *CategoryRepo) FindCategoryById(Id int64) (*model.Category, error) {
	if category, ok := r.categories[Id]; ok {
		return category, nil //tìm được
	} else {
		return nil, errors.New("category not found")
	}
}

func (r *CategoryRepo) DeleteCategoryById(Id int64) error {
	if _, ok := r.categories[Id]; ok {
		delete(r.categories, Id)
		return nil
	} else {
		return errors.New("category not found")
	}
}

func (r *CategoryRepo) UpdateCategory(category *model.Category) error {
	if _, ok := r.categories[category.Id]; ok {
		r.categories[category.Id] = category
		return nil //tìm được
	} else {
		return errors.New("category not found")
	}
}

func (r *CategoryRepo) Upsert(category *model.Category) int64 {
	if _, ok := r.categories[category.Id]; ok {
		r.categories[category.Id] = category //tìm thấy thì update
		return category.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewCategory(category)
	}
}
