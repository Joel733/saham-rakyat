package mysql

import (
	"saham-rakyat/main/entity"
	"saham-rakyat/main/pkg/pagination"
	"saham-rakyat/main/repository"

	"gorm.io/gorm"
)

type userTable struct {
	db *gorm.DB
}

func NewUserTable(db *gorm.DB) repository.UsersRepository{
	return &userTable{
		db: db,
	}
}

func (t *userTable) Create(e *entity.Users) (*entity.Users, error) {
	tx := t.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (t *userTable) FindAllBy(page, size int) ([]*entity.Users, error) {
	limit, offset := pagination.GetLimitOffset(page, size)

	var users []*entity.Users

	res := t.db.Model(&users).
		Limit(limit).
		Offset(offset).
		Scan(&users)

	if err := res.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (t *userTable) Detail(id int) (*entity.Users ,error){
	
	var userDetail entity.Users

	res := t.db.Limit(1).Find(&userDetail, id)

	if err := res.Error; err != nil {
		return nil, err
	}

	return &userDetail, nil
}

func (t *userTable) Update(e *entity.Users) (*entity.Users, error){
	
	var updatedUser entity.Users

	res := t.db.Save(&e)

	if err := res.Error; err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (t *userTable) Delete( e *entity.Users) error{

	return t.db.Delete(&e).Error
}
