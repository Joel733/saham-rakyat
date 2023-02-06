package request

import "time"

type OrderHistoriesRequest struct {
	ID           int       `json:"id"`
	UserId       int       `json:"user_id"`
	OrderItemId  int   	   `json:"first_order"`
	Descriptions string    `json:"descriptions"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FindAllOrderHistoriesItems struct{
	Page int `json:"page" query:"page"`
	Size int `json:"size" query:"size"`
}
