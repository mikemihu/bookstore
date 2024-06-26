package usecase

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/constant"
	"gotu-bookstore/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BookUC struct {
	logger   *zap.Logger
	bookRepo internal.BookRepo
}

func NewBookUC(
	logger *zap.Logger,
	bookRepo internal.BookRepo,
) internal.BookUC {
	return &BookUC{
		logger:   logger,
		bookRepo: bookRepo,
	}
}

func (p *BookUC) GetList(ctx context.Context, req entity.BookGetListRequest) ([]entity.BookResponse, error) {
	filter := entity.BookFilter{
		Search: req.Search,
		IDs:    req.IDs,
	}
	books, err := p.bookRepo.Get(ctx, filter)
	if err != nil {
		if !errors.Is(err, constant.ErrNotFound) {
			p.logger.Error("failed bookRepo.Get", zap.Error(err),
				zap.Any("filter", filter))
		}
		return nil, err
	}

	resp := make([]entity.BookResponse, len(books))
	for i := range books {
		resp[i] = books[i].ToResponse()
	}
	return resp, nil
}

func (p *BookUC) Get(ctx context.Context, id uuid.UUID) (entity.BookResponse, error) {
	filter := entity.BookFilter{
		Book: entity.Book{
			BaseModel: entity.BaseModel{ID: id},
		},
	}
	books, err := p.bookRepo.Get(ctx, filter)
	if err != nil {
		if !errors.Is(err, constant.ErrNotFound) {
			p.logger.Error("failed bookRepo.Get", zap.Error(err),
				zap.Any("filter", filter))
		}
		return entity.BookResponse{}, err
	}
	book := books[0]

	return book.ToResponse(), nil
}
