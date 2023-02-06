package repository

import "saham-rakyat/main/entity"

type UsersRepository interface {
	Create(e *entity.Users) (*entity.Users ,error)
	FindAllBy(page, size int) ([] *entity.Users,error)
	Detail(id int)(*entity.Users, error)
	Update(e *entity.Users) (*entity.Users,error)
	Delete(e *entity.Users) (error) 
	
}
