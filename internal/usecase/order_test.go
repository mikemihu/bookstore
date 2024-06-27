package usecase

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/constant"
	"gotu-bookstore/internal/contexts"
	"gotu-bookstore/internal/entity"
	"gotu-bookstore/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestOrderUC_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyUserID := uuid.New()
	dummyOrderID1 := uuid.New()
	dummyOrderID2 := uuid.New()
	dummyBookID1 := uuid.New()
	dummyBookID2 := uuid.New()

	type fields struct {
		orderRepo internal.OrderRepo
		bookRepo  internal.BookRepo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.OrderResponse
		wantErr error
	}{
		{
			name: "success - with some records",
			fields: fields{
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					filter := entity.OrderFilter{
						Order: entity.Order{
							UserID: dummyUserID,
						},
					}
					orders := []entity.Order{
						{
							BaseModel: entity.BaseModel{ID: dummyOrderID1},
							UserID:    dummyUserID,
							Items: []entity.OrderItem{
								{BookID: dummyBookID1, Qty: 1},
							},
						},
						{
							BaseModel: entity.BaseModel{ID: dummyOrderID2},
							UserID:    dummyUserID,
							Items: []entity.OrderItem{
								{BookID: dummyBookID2, Qty: 2},
							},
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(orders, nil)
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
			},
			want: []entity.OrderResponse{
				{
					ID:     dummyOrderID1,
					UserID: dummyUserID,
					Items: []entity.OrderItemResponse{
						{BookID: dummyBookID1, Qty: 1},
					},
				},
				{
					ID:     dummyOrderID2,
					UserID: dummyUserID,
					Items: []entity.OrderItemResponse{
						{BookID: dummyBookID2, Qty: 2},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "error - no order yet",
			fields: fields{
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					filter := entity.OrderFilter{
						Order: entity.Order{
							UserID: dummyUserID,
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, constant.ErrNotFound)
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
			},
			want:    nil,
			wantErr: constant.ErrNoOrderYet,
		},
		{
			name: "error - from order repo get",
			fields: fields{
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					filter := entity.OrderFilter{
						Order: entity.Order{
							UserID: dummyUserID,
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, errors.New("some error"))
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
			},
			want:    nil,
			wantErr: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &OrderUC{
				logger:    zap.NewNop(),
				orderRepo: tt.fields.orderRepo,
				bookRepo:  tt.fields.bookRepo,
			}
			got, err := p.GetList(tt.args.ctx)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOrderUC_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyUserID := uuid.New()
	dummyOrderID := uuid.New()
	dummyBookID1 := uuid.New()
	dummyBookID2 := uuid.New()

	type fields struct {
		orderRepo internal.OrderRepo
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.OrderResponse
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					filter := entity.OrderFilter{
						Order: entity.Order{
							BaseModel: entity.BaseModel{ID: dummyOrderID},
							UserID:    dummyUserID,
						},
						PreloadDetail: true,
					}
					orders := []entity.Order{
						{
							BaseModel: entity.BaseModel{ID: dummyOrderID},
							UserID:    dummyUserID,
							Items: []entity.OrderItem{
								{BookID: dummyBookID1, Qty: 1},
								{BookID: dummyBookID2, Qty: 2},
							},
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(orders, nil)
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
				id: dummyOrderID,
			},
			want: entity.OrderResponse{
				ID:     dummyOrderID,
				UserID: dummyUserID,
				Items: []entity.OrderItemResponse{
					{BookID: dummyBookID1, Qty: 1},
					{BookID: dummyBookID2, Qty: 2},
				},
			},
			wantErr: nil,
		},
		{
			name: "error - order not found",
			fields: fields{
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					filter := entity.OrderFilter{
						Order: entity.Order{
							BaseModel: entity.BaseModel{ID: dummyOrderID},
							UserID:    dummyUserID,
						},
						PreloadDetail: true,
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, constant.ErrNotFound)
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
				id: dummyOrderID,
			},
			want:    entity.OrderResponse{},
			wantErr: constant.ErrNotFound,
		},
		{
			name: "error - other error from order repo get",
			fields: fields{
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					filter := entity.OrderFilter{
						Order: entity.Order{
							BaseModel: entity.BaseModel{ID: dummyOrderID},
							UserID:    dummyUserID,
						},
						PreloadDetail: true,
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, errors.New("some error"))
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
				id: dummyOrderID,
			},
			want:    entity.OrderResponse{},
			wantErr: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &OrderUC{
				logger:    zap.NewNop(),
				orderRepo: tt.fields.orderRepo,
			}
			got, err := p.Get(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOrderUC_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyUserID := uuid.New()
	dummyOrderID := uuid.New()
	dummyBookID1 := uuid.New()
	dummyBookID2 := uuid.New()

	type fields struct {
		orderRepo internal.OrderRepo
		bookRepo  internal.BookRepo
	}
	type args struct {
		ctx context.Context
		req entity.OrderCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uuid.UUID
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter1 := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyBookID1},
						},
					}
					filter2 := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyBookID2},
						},
					}
					books1 := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyBookID1}, ISBN: "1234567890123", Price: 10},
					}
					books2 := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyBookID2}, ISBN: "2345678901234", Price: 15},
					}
					mock.EXPECT().Get(gomock.Any(), filter1).Return(books1, nil)
					mock.EXPECT().Get(gomock.Any(), filter2).Return(books2, nil)
					return mock
				}(),
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					order := entity.Order{
						UserID: dummyUserID,
						Items: []entity.OrderItem{
							{BookID: dummyBookID1, Qty: 1, Price: 10},
							{BookID: dummyBookID2, Qty: 2, Price: 15},
						},
						TotalQty:   3,
						TotalPrice: 40,
					}
					mock.EXPECT().Store(gomock.Any(), order).Return(dummyOrderID, nil)
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
				req: entity.OrderCreateRequest{
					Items: []struct {
						BookID uuid.UUID `json:"book_id"`
						Qty    int       `json:"qty"`
					}{
						{BookID: dummyBookID1, Qty: 1},
						{BookID: dummyBookID2, Qty: 2},
					},
				},
			},
			want:    dummyOrderID,
			wantErr: nil,
		},
		{
			name: "error - from book repo get",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter1 := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyBookID1},
						},
					}
					filter2 := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyBookID2},
						},
					}
					books1 := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyBookID1}, ISBN: "1234567890123", Price: 10},
					}
					mock.EXPECT().Get(gomock.Any(), filter1).Return(books1, nil)
					mock.EXPECT().Get(gomock.Any(), filter2).Return(nil, errors.New("some error from book repo"))
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
				req: entity.OrderCreateRequest{
					Items: []struct {
						BookID uuid.UUID `json:"book_id"`
						Qty    int       `json:"qty"`
					}{
						{BookID: dummyBookID1, Qty: 1},
						{BookID: dummyBookID2, Qty: 2},
					},
				},
			},
			want:    uuid.Nil,
			wantErr: errors.New("some error from book repo"),
		},
		{
			name: "error - from order repo store",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter1 := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyBookID1},
						},
					}
					filter2 := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyBookID2},
						},
					}
					books1 := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyBookID1}, ISBN: "1234567890123", Price: 10},
					}
					books2 := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyBookID2}, ISBN: "2345678901234", Price: 15},
					}
					mock.EXPECT().Get(gomock.Any(), filter1).Return(books1, nil)
					mock.EXPECT().Get(gomock.Any(), filter2).Return(books2, nil)
					return mock
				}(),
				orderRepo: func() *mocks.MockOrderRepo {
					mock := mocks.NewMockOrderRepo(ctrl)
					order := entity.Order{
						UserID: dummyUserID,
						Items: []entity.OrderItem{
							{BookID: dummyBookID1, Qty: 1, Price: 10},
							{BookID: dummyBookID2, Qty: 2, Price: 15},
						},
						TotalQty:   3,
						TotalPrice: 40,
					}
					mock.EXPECT().Store(gomock.Any(), order).Return(uuid.Nil, errors.New("some error from order repo"))
					return mock
				}(),
			},
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyUserID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
				req: entity.OrderCreateRequest{
					Items: []struct {
						BookID uuid.UUID `json:"book_id"`
						Qty    int       `json:"qty"`
					}{
						{BookID: dummyBookID1, Qty: 1},
						{BookID: dummyBookID2, Qty: 2},
					},
				},
			},
			want:    uuid.Nil,
			wantErr: errors.New("some error from order repo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &OrderUC{
				logger:    zap.NewNop(),
				orderRepo: tt.fields.orderRepo,
				bookRepo:  tt.fields.bookRepo,
			}
			got, err := p.Create(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
