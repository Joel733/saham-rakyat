package request

import "time"

type UsersRequest struct {
	ID           int       `json:"id"`
	FullName     string    `json:"full_name"`
	FirstOrder   string    `json:"first_order"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type FindAllUsersRequest struct{
	Page int `json:"page" query:"page"`
	Size int `json:"size" query:"size"`
}
