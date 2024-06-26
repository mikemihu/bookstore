package repository

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/constant"
	"gotu-bookstore/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepo struct {
	cfg *config.Cfg
	db  *gorm.DB
}

func NewOrderRepo(
	cfg *config.Cfg,
	db *gorm.DB,
) internal.OrderRepo {
	return &OrderRepo{
		cfg: cfg,
		db:  db,
	}
}

func (p *OrderRepo) Get(ctx context.Context, filter entity.OrderFilter) ([]entity.Order, error) {
	tx := p.applyFilter(p.db.WithContext(ctx), filter)

	// single / multiple records
	var orders []entity.Order
	err := tx.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, constant.ErrNotFound
	}
	return orders, nil
}

func (p *OrderRepo) applyFilter(tx *gorm.DB, filter entity.OrderFilter) *gorm.DB {
	if filter.ID != uuid.Nil {
		tx = tx.Where("id = ?", filter.ID)
	}
	if filter.UserID != uuid.Nil {
		tx = tx.Where("user_id = ?", filter.UserID)
	}
	if filter.PreloadDetail {
		tx = tx.Preload("Items").
			Preload("Items.Book")
	}
	return tx
}

func (p *OrderRepo) Store(ctx context.Context, order entity.Order) (uuid.UUID, error) {
	tx := p.db.WithContext(ctx)

	if order.ID == uuid.Nil {
		// create new record

		// generate uuid
		order.ID = uuid.New()
		err := tx.Omit("Items.ID").Create(&order).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return uuid.Nil, constant.ErrDuplicateRecord
		}
		if err != nil {
			return uuid.Nil, err
		}

		return order.ID, nil
	}

	// update existing record
	tx = tx.Model(&order).Updates(&order)
	err := tx.Error
	if err != nil {
		return uuid.Nil, err
	}
	if tx.RowsAffected == 0 {
		return uuid.Nil, constant.ErrNotFound
	}

	return order.ID, nil
}
