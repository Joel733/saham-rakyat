package service

import (
	"fmt"
	"saham-rakyat/main/entity"
	"saham-rakyat/main/repository"
	"saham-rakyat/main/request"
	"time"
) 

type ordersItemRepo struct {
	repo repository.OrdersItemsRepository
}

type OrdersItemService interface {
	Create(req request.OrdersItemRequest) (interface{},error)
	FindAll(req request.OrdersItemRequest) (interface{},error)
	Detail(id int)(interface{},error)
	Update(req request.OrdersItemRequest) (interface{},error)
	Delete(req request.OrdersItemRequest) (interface{},error)
}

func (u *ordersItemRepo) Create(req request.OrdersItemRequest) (interface{}, error) {
	t := entity.OrdersItem{
		ID:     req.ID,
		Name:   req.Name,
		Price:     req.Price,
		ExpiredAt:      req.ExpiredAt,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		DeletedAt : req.DeletedAt,
	}

	team, err := u.repo.Create(&t)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r *ordersItemRepo) FindAll(req request.FindAllOrderItems) (interface{}, error) {
	teams, err := r.repo.FindAllBy(req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func(r *ordersItemRepo) Detail(id int) (interface{},error){
	order, err := r.repo.Detail(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *ordersItemRepo) Update(req request.OrdersItemRequest) (interface{}, error){

	order, err := r.repo.Detail(req.ID)

	if err != nil {
		return nil, err
	}
	
	order.Name = req.Name
	order.Price = req.Price
	order.ExpiredAt = req.ExpiredAt
	order.UpdatedAt = time.Now()

	return order, nil
}

func (r *ordersItemRepo) Delete(req request.OrdersItemRequest) (error){

	order, err := r.repo.Detail(req.ID)

	if err != nil {
		return fmt.Errorf("failed to find order histories : %v",err)
	}

	err = r.repo.Delete(order)

	return err
}