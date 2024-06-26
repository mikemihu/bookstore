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

type BookRepo struct {
	cfg *config.Cfg
	db  *gorm.DB
}

func NewBookRepo(
	cfg *config.Cfg,
	db *gorm.DB,
) internal.BookRepo {
	return &BookRepo{
		cfg: cfg,
		db:  db,
	}
}

func (p *BookRepo) Get(ctx context.Context, filter entity.BookFilter) ([]entity.Book, error) {
	tx := p.applyFilter(p.db.WithContext(ctx), filter)

	// single / multiple records
	var books []entity.Book
	err := tx.Find(&books).Error
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, constant.ErrNotFound
	}
	return books, nil
}

func (p *BookRepo) applyFilter(tx *gorm.DB, filter entity.BookFilter) *gorm.DB {
	if filter.ID != uuid.Nil {
		tx = tx.Where("id = ?", filter.ID)
	}
	if len(filter.IDs) != 0 {
		tx = tx.Where("id IN ?", filter.IDs)
	}
	if len(filter.Search) != 0 {
		keyword := "%" + filter.Search + "%"
		tx = tx.Where("author ILIKE ? OR title ILIKE ? OR subtitle ILIKE ?", keyword, keyword, keyword)
	}
	return tx
}

func (p *BookRepo) Store(ctx context.Context, book entity.Book) (uuid.UUID, error) {
	tx := p.db.WithContext(ctx)

	// create new record
	if book.ID == uuid.Nil {
		book.ID = uuid.New()
		err := tx.Create(&book).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return uuid.Nil, constant.ErrDuplicateRecord
		}
		if err != nil {
			return uuid.Nil, err
		}

		return book.ID, nil
	}

	// update existing record
	tx = tx.Model(&book).Updates(&book)
	err := tx.Error
	if err != nil {
		return uuid.Nil, err
	}
	if tx.RowsAffected == 0 {
		return uuid.Nil, constant.ErrNotFound
	}

	return book.ID, nil
}
