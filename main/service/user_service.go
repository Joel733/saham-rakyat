package service

import (
	"fmt"
	"saham-rakyat/main/entity"
	"saham-rakyat/main/repository"
	"saham-rakyat/main/request"
	"time"
) 

type userRepository struct {
	repo repository.UsersRepository
}

type UserService interface {
	Create(req request.UsersRequest) (interface{},error)
	FindAll(req request.UsersRequest) (interface{},error)
	Detail(id int)(interface{},error)
	Update(req request.UsersRequest) (interface{},error)
	Delete(req request.UsersRequest) (interface{},error)
}

func (u *userRepository) Create(req request.UsersRequest) (interface{}, error) {
	t := entity.Users{
		ID:     req.ID,
		FullName: 	 req.FullName,
		FirstOrder:     req.FirstOrder,
		CreatedAt: 		req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		DeletedAt : req.DeletedAt,
	}

	team, err := u.repo.Create(&t)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r *userRepository) FindAll(req request.FindAllUsersRequest) (interface{}, error) {
	teams, err := r.repo.FindAllBy(req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func(r *userRepository) Detail (id int) (interface{},error){
	order, err := r.repo.Detail(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *userRepository) Update(req request.UsersRequest) (interface{}, error){

	order, err := r.repo.Detail(req.ID)

	if err != nil {
		return nil, err
	}
	
	order.FullName = req.FullName
	order.UpdatedAt = time.Now()	

	return order, nil
}

func (r *userRepository) Delete(req request.UsersRequest) (error){

	order, err := r.repo.Detail(req.ID)

	if err != nil {
		return fmt.Errorf("failed to find user : %v",err)
	}

	err = r.repo.Delete(order)

	return err
}