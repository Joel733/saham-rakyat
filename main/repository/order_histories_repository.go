package repository

import "saham-rakyat/main/entity"

type OrderHistoriesRepository interface {
	Create(e *entity.OrderHistories) (*entity.OrderHistories ,error)
	FindAllBy(page, size int) ([] *entity.OrderHistories,error)
	Detail(id int)(*entity.OrderHistories, error)
	Update(e *entity.OrderHistories) (*entity.OrderHistories,error)
	Delete(e *entity.OrderHistories) (error) 
}