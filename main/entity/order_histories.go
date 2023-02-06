package entity

import "time"

type OrderHistories struct {
	ID           int       `json:"id"`
	UserId       int    `json:"full_name"`
	OrderItemId  int    `json:"first_order"`
	Descriptions string    `json:"descriptions"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}