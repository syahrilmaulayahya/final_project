package requests

import "final_project/business/users"

type UserReview struct {
	Review    string  `json:"review"`
	Rating    float32 `json:"rating"`
	ProductID int     `json:"productid"`
}

func (review *UserReview) ToDomain() users.Review_RatingDomain {
	return users.Review_RatingDomain{
		Review:    review.Review,
		Rating:    review.Rating,
		ProductID: review.ProductID,
	}
}
