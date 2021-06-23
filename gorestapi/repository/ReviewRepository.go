package repository

import (
	"errors"
	"fmt"

	"github.com/ocg.com/go/01/lab/gorestapi/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64

	observerList []Observer
}

var Reviews ReviewRepo //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
	Reviews.InitData("sql:45312")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *ProductRepo
func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	review.Id = nextID
	r.reviews[nextID] = review
	return nextID
}

func (r *ReviewRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewReview(&model.Review{
		ProductId: 1,
		Comment:   "good",
		Rating:    4,
	})

	r.CreateNewReview(&model.Review{
		ProductId: 2,
		Comment:   "bad",
		Rating:    2,
	})

	r.CreateNewReview(&model.Review{
		ProductId: 2,
		Comment:   "bad",
		Rating:    2,
	})

	r.CreateNewReview(&model.Review{
		ProductId: 1,
		Comment:   "good",
		Rating:    4,
	})
}

func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepo) FindReviewById(Id int64) (*model.Review, error) {
	if review, ok := r.reviews[Id]; ok {
		return review, nil //tìm được
	} else {
		return nil, errors.New("review not found")
	}
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) UpdateReview(review *model.Review) error {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return nil //tìm được
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) Upsert(review *model.Review) int64 {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return review.Id
	} else {
		return r.CreateNewReview(review)
	}
}

//--------------
/*
Khi một rewview thay đổi thì chạy làm hàm này để tính lại average product rating
*/
func (r *ReviewRepo) ComputeProductRatingWhenAReviewChage(review *model.Review) (productId int64, averageRating float32) {
	productId = review.ProductId
	return productId, r.ComputeProductRating(productId)
}

func (r *ReviewRepo) ComputeProductRating(productId int64) float32 {
	count := 0
	sum := 0
	for _, review := range r.reviews {
		if review.ProductId == productId {
			count++
			sum += review.Rating
		}
	}
	return float32(sum / count)
}

//-------------- Impement interface Publisher ----
func (r *ReviewRepo) RegisterObserver(o Observer) {
	r.observerList = append(r.observerList, o)
}

func (r *ReviewRepo) RemoveObserver(o Observer) {
	found := false
	i := 0
	for ; i < len(r.observerList); i++ {
		if r.observerList[i] == o {
			found = true
			break
		}
	}
	if found {
		r.observerList = append(r.observerList[:i], r.observerList[i+1:]...)
	}
}

func (r *ReviewRepo) NotifyObserver(id int64) {
	review := r.reviews[id]
	productId, averageRating := r.ComputeProductRatingWhenAReviewChage(review)

	for _, observer := range r.observerList {
		observer.Update(productId, averageRating)
	}
}
