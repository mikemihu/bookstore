package usecase

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/constant"
	"gotu-bookstore/internal/contexts"
	"gotu-bookstore/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type OrderUC struct {
	logger    *zap.Logger
	orderRepo internal.OrderRepo
	bookRepo  internal.BookRepo
}

func NewOrderUC(
	logger *zap.Logger,
	orderRepo internal.OrderRepo,
	bookRepo internal.BookRepo,
) internal.OrderUC {
	return &OrderUC{
		logger:    logger,
		orderRepo: orderRepo,
		bookRepo:  bookRepo,
	}
}

func (p *OrderUC) GetList(ctx context.Context) ([]entity.OrderResponse, error) {
	userID := contexts.GetUser(ctx).ID
	filter := entity.OrderFilter{
		Order: entity.Order{
			UserID: userID,
		},
	}
	orders, err := p.orderRepo.Get(ctx, filter)
	if errors.Is(err, constant.ErrNotFound) {
		return nil, constant.ErrNoOrderYet
	}
	if err != nil {
		p.logger.Error("failed orderRepo.Get", zap.Error(err),
			zap.String("user_id", userID.String()))
		return nil, err
	}

	resp := make([]entity.OrderResponse, len(orders))
	for i := range orders {
		resp[i] = orders[i].ToResponse()
	}
	return resp, nil
}

func (p *OrderUC) Get(ctx context.Context, id uuid.UUID) (entity.OrderResponse, error) {
	userID := contexts.GetUser(ctx).ID
	filter := entity.OrderFilter{
		Order: entity.Order{
			BaseModel: entity.BaseModel{ID: id},
			UserID:    userID, // to make sure user's can only get his own orders
		},
		PreloadDetail: true,
	}
	orders, err := p.orderRepo.Get(ctx, filter)
	if err != nil {
		if !errors.Is(err, constant.ErrNotFound) {
			p.logger.Error("failed orderRepo.Get", zap.Error(err),
				zap.String("user_id", userID.String()),
				zap.String("order_id", id.String()))
		}
		return entity.OrderResponse{}, err
	}
	order := orders[0]
	return order.ToResponse(), nil
}

func (p *OrderUC) Create(ctx context.Context, req entity.OrderCreateRequest) (uuid.UUID, error) {
	userID := contexts.GetUser(ctx).ID
	order := entity.Order{
		UserID: userID,
		Items:  make([]entity.OrderItem, len(req.Items)),
	}

	for i := range req.Items {
		item := &req.Items[i]

		// check if book exists
		bookFilter := entity.BookFilter{
			Book: entity.Book{
				BaseModel: entity.BaseModel{ID: item.BookID},
			},
		}
		books, err := p.bookRepo.Get(ctx, bookFilter)
		if err != nil {
			p.logger.Error("failed bookRepo.Get", zap.Error(err),
				zap.String("book_id", item.BookID.String()))
			return uuid.Nil, err
		}
		book := books[0]

		// fill order's item
		orderItem := entity.OrderItem{
			BookID: item.BookID,
			Qty:    item.Qty,
			Price:  book.Price,
		}
		order.Items[i] = orderItem

		// accumulate total qty & price
		order.TotalQty = order.TotalQty + item.Qty
		order.TotalPrice = order.TotalPrice + (float64(item.Qty) * book.Price)
	}

	// store order
	id, err := p.orderRepo.Store(ctx, order)
	if err != nil {
		p.logger.Error("failed orderRepo.Store", zap.Error(err),
			zap.String("user_id", userID.String()),
			zap.Any("order", order))
		return uuid.Nil, err
	}

	return id, nil
}
