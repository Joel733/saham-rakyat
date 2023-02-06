package mysql

import (
	"saham-rakyat/main/entity"
	"saham-rakyat/main/pkg/pagination"
	"saham-rakyat/main/repository"

	"gorm.io/gorm"
)

type orderItemsTable struct {
	db *gorm.DB
}

func NewOrderItemsTable(db *gorm.DB) repository.OrdersItemsRepository {
	return &orderItemsTable{
		db: db,
	}
}

func (t *orderItemsTable) Create(e *entity.OrdersItem) (*entity.OrdersItem, error) {
	tx := t.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (t *orderItemsTable) FindAllBy(page, size int) ([]*entity.OrdersItem, error) {
	limit, offset := pagination.GetLimitOffset(page, size)

	var orders []*entity.OrdersItem

	res := t.db.Model(&orders).
		Limit(limit).
		Offset(offset).
		Scan(&orders)

	if err := res.Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (t *orderItemsTable) Detail(id int) (*entity.OrdersItem ,error){
	
	var orderDetail entity.OrdersItem

	res := t.db.Limit(1).Find(&orderDetail, id)

	if err := res.Error; err != nil {
		return nil, err
	}

	return &orderDetail, nil
}

func (t *orderItemsTable) Update(e *entity.OrdersItem) (*entity.OrdersItem, error){
	
	var updatedOrder entity.OrdersItem

	res := t.db.Save(&e)

	if err := res.Error; err != nil {
		return nil, err
	}

	return &updatedOrder, nil
}

func (t *orderItemsTable) Delete( e *entity.OrdersItem) error{

	return t.db.Delete(&e).Error
}