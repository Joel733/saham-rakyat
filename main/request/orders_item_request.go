package request

import "time"

type OrdersItemRequest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	ExpiredAt int       `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindAllOrderItems struct{
	Page int `json:"page" query:"page"`
	Size int `json:"size" query:"size"`
}

