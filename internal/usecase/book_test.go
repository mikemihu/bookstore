package usecase

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/entity"
	"gotu-bookstore/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestBookUC_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyID1 := uuid.New()
	dummyID2 := uuid.New()

	type fields struct {
		bookRepo internal.BookRepo
	}
	type args struct {
		req entity.BookGetListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.BookResponse
		wantErr error
	}{
		{
			name: "success - with all query",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter := entity.BookFilter{
						Search: "test",
						IDs: []uuid.UUID{
							dummyID1,
							dummyID2,
						},
					}
					books := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyID1}, ISBN: "1234567890123"},
						{BaseModel: entity.BaseModel{ID: dummyID2}, ISBN: "2345678901234"},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(books, nil)
					return mock
				}(),
			},
			args: args{
				req: entity.BookGetListRequest{
					Search: "test",
					IDs: []uuid.UUID{
						dummyID1,
						dummyID2,
					},
				},
			},
			want: []entity.BookResponse{
				{ID: dummyID1, ISBN: "1234567890123"},
				{ID: dummyID2, ISBN: "2345678901234"},
			},
		},
		{
			name: "success - without any query",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter := entity.BookFilter{
						Search: "",
						IDs:    []uuid.UUID{},
					}
					books := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyID1}, ISBN: "1234567890123"},
						{BaseModel: entity.BaseModel{ID: dummyID2}, ISBN: "2345678901234"},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(books, nil)
					return mock
				}(),
			},
			args: args{
				req: entity.BookGetListRequest{
					Search: "",
					IDs:    []uuid.UUID{},
				},
			},
			want: []entity.BookResponse{
				{ID: dummyID1, ISBN: "1234567890123"},
				{ID: dummyID2, ISBN: "2345678901234"},
			},
		},
		{
			name: "error - from book repo get",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter := entity.BookFilter{
						Search: "test",
						IDs: []uuid.UUID{
							dummyID1,
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, errors.New("some error"))
					return mock
				}(),
			},
			args: args{
				req: entity.BookGetListRequest{
					Search: "test",
					IDs:    []uuid.UUID{dummyID1},
				},
			},
			want:    nil,
			wantErr: errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &BookUC{
				logger:   zap.NewNop(),
				bookRepo: tt.fields.bookRepo,
			}
			got, err := p.GetList(context.Background(), tt.args.req)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBookUC_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyID1 := uuid.New()

	type fields struct {
		bookRepo internal.BookRepo
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.BookResponse
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyID1},
						},
					}
					books := []entity.Book{
						{BaseModel: entity.BaseModel{ID: dummyID1}, ISBN: "1234567890123"},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(books, nil)
					return mock
				}(),
			},
			args: args{
				id: dummyID1,
			},
			want: entity.BookResponse{
				ID: dummyID1, ISBN: "1234567890123",
			},
			wantErr: nil,
		},
		{
			name: "error - from book repo get",
			fields: fields{
				bookRepo: func() *mocks.MockBookRepo {
					mock := mocks.NewMockBookRepo(ctrl)
					filter := entity.BookFilter{
						Book: entity.Book{
							BaseModel: entity.BaseModel{ID: dummyID1},
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, errors.New("some error"))
					return mock
				}(),
			},
			args: args{
				id: dummyID1,
			},
			want:    entity.BookResponse{},
			wantErr: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &BookUC{
				logger:   zap.NewNop(),
				bookRepo: tt.fields.bookRepo,
			}
			got, err := p.Get(context.Background(), tt.args.id)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
