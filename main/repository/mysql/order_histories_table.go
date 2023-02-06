package mysql

import (
	"saham-rakyat/main/entity"
	"saham-rakyat/main/pkg/pagination"
	"saham-rakyat/main/repository"

	"gorm.io/gorm"
)

type orderHistoriesTable struct {
	db *gorm.DB
}

func NewOrderHistoriesTable(db *gorm.DB) repository.OrderHistoriesRepository {
	return &orderHistoriesTable{
		db: db,
	}
}

func (t *orderHistoriesTable) Create(e *entity.OrderHistories) (*entity.OrderHistories, error) {
	tx := t.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (t *orderHistoriesTable) FindAllBy(page, size int) ([]*entity.OrderHistories, error){
	limit, offset := pagination.GetLimitOffset(page, size)

	var orderHistories []*entity.OrderHistories

	res := t.db.Model(&orderHistories).
		Limit(limit).
		Offset(offset).
		Scan(&orderHistories)

	if err := res.Error; err != nil {
		return nil, err
	}

	return orderHistories, nil
}

func (t *orderHistoriesTable) Detail(id int) (*entity.OrderHistories ,error){
	
	var orderDetail entity.OrderHistories

	res := t.db.Limit(1).Find(&orderDetail, id)

	if err := res.Error; err != nil {
		return nil, err
	}

	return &orderDetail, nil
}

func (t *orderHistoriesTable) Update(e *entity.OrderHistories) (*entity.OrderHistories, error){
	
	var updatedOrder entity.OrderHistories

	res := t.db.Save(&e)

	if err := res.Error; err != nil {
		return nil, err
	}

	return &updatedOrder, nil
}

func (t *orderHistoriesTable) Delete( e *entity.OrderHistories) error{

	return t.db.Delete(&e).Error
}
