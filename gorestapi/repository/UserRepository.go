package repository

import (
	"errors"
	"fmt"

	"github.com/ocg.com/go/01/lab/gorestapi/model"
)

type UserRepo struct {
	users  map[int64]*model.User
	autoID int64 //đây là biến đếm tự tăng gán giá trị cho id của user
}

var Users UserRepo

func init() {
	Users = UserRepo{autoID: 0}
	Users.users = make(map[int64]*model.User)
	Users.InitData("sql:45312")
}

func (r *UserRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *UserRepo) CreateNewUser(user *model.User) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	user.Id = nextID
	r.users[nextID] = user //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *UserRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewUser(&model.User{
		FirstName:  "Administrator",
		LastName:   "",
		UserName:   "Admintrator",
		Email:      "admin@gmail.com",
		Password:   "admin",
		Avatar:     "https://robohash.org/eaquequasincidunt.png?size=50x50&set=set1",
		Gender:     "Genderfluid",
		Phone:      "933-658-1213",
		Birthday:   "1994-03-23",
		Status:     true,
		CreatedAt:  1609483221000,
		ModifiedAt: 1609483221000})

	r.CreateNewUser(&model.User{
		FirstName:  "Client 1",
		LastName:   "",
		UserName:   "client1",
		Email:      "client1@gmail.com",
		Password:   "client",
		Avatar:     "https://robohash.org/accusantiumminimamagni.png?size=50x50&set=set1",
		Gender:     "Male",
		Phone:      "510-449-7332",
		Birthday:   "2002-03-11",
		Status:     false,
		CreatedAt:  1617440961000,
		ModifiedAt: 1618301961000})

	r.CreateNewUser(&model.User{
		FirstName:  "Client 2",
		LastName:   "",
		UserName:   "client2",
		Email:      "client2@gmail.com",
		Password:   "kjU6qK1Bm",
		Avatar:     "https://robohash.org/voluptatemdebitiset.png?size=50x50&set=set1",
		Gender:     "Female",
		Phone:      "676-983-4977",
		Birthday:   "1997-09-29",
		Status:     false,
		CreatedAt:  1615745961000,
		ModifiedAt: 1615976361000})
}

func (r *UserRepo) GetAllUsers() map[int64]*model.User {
	return r.users
}

func (r *UserRepo) FindUserById(Id int64) (*model.User, error) {
	if user, ok := r.users[Id]; ok {
		return user, nil //tìm được
	} else {
		return nil, errors.New("user not found")
	}
}

func (r *UserRepo) DeleteUserById(Id int64) error {
	if _, ok := r.users[Id]; ok {
		delete(r.users, Id)
		return nil
	} else {
		return errors.New("user not found")
	}
}

func (r *UserRepo) UpdateUser(user *model.User) error {
	if _, ok := r.users[user.Id]; ok {
		r.users[user.Id] = user
		return nil //tìm được
	} else {
		return errors.New("user not found")
	}
}

func (r *UserRepo) Upsert(user *model.User) int64 {
	if _, ok := r.users[user.Id]; ok {
		r.users[user.Id] = user //tìm thấy thì update
		return user.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewUser(user)
	}
}
