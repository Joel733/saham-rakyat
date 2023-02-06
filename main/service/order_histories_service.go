package service

import (
	"fmt"
	"saham-rakyat/main/entity"
	"saham-rakyat/main/repository"
	"saham-rakyat/main/request"
) 

type orderHistoriesRepo struct {
	repo repository.OrderHistoriesRepository
}

type OrdersHistoriesService interface {
	Create(req request.OrderHistoriesRequest) (interface{},error)
	FindAll(req request.FindAllOrderHistoriesItems) (interface{},error)
	Detail(id int)(interface{},error)
	Update(req request.OrderHistoriesRequest) (interface{},error)
	Delete(req request.OrderHistoriesRequest) (interface{},error)

}

func NewOrderHistoriesService(repo repository.OrderHistoriesRepository) OrdersHistoriesService{
	return &orderHistoriesRepo{repo}
}

func (r *orderHistoriesRepo) Create(req request.OrderHistoriesRequest) (interface{}, error) {
	t := entity.OrderHistories{
		ID:     req.ID,
		UserId:  req.UserId,
		OrderItemId:     req.OrderItemId,
		Descriptions:  req.Descriptions,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		DeletedAt : req.DeletedAt,
	}

	team, err := r.repo.Create(&t)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r *orderHistoriesRepo) FindAll(req request.FindAllOrderHistoriesItems) (interface{}, error) {
	orders, err := r.repo.FindAllBy(req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func(r *orderHistoriesRepo) Detail(id int) (interface{},error){
	order, err := r.repo.Detail(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *orderHistoriesRepo) Update(req request.OrderHistoriesRequest) (interface{}, error){

	order, err := r.repo.Detail(req.ID)

	if err != nil {
		return nil, err
	}

	order.Descriptions = req.Descriptions
	order.UpdatedAt = req.UpdatedAt

	return order, nil
}

func (r *orderHistoriesRepo) Delete(req request.OrderHistoriesRequest) (interface{},error){

	order, err := r.repo.Detail(req.ID)

	if err != nil {
		return nil,fmt.Errorf("failed to find order histories : %v",err)
	}

	err = r.repo.Delete(order)

	rsp := map[string]string{"status" : "OK", "message" : "Delete success"}

	return rsp, err
}