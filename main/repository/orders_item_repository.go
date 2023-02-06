package repository

import "saham-rakyat/main/entity"

type OrdersItemsRepository interface {
	Create(e *entity.OrdersItem) (*entity.OrdersItem, error)
	FindAllBy(page, size int) ([] *entity.OrdersItem,error)
	Detail(id int)(*entity.OrdersItem, error)
	Update(e *entity.OrdersItem) (*entity.OrdersItem,error)
	Delete(e *entity.OrdersItem) (error) 
}

