package usecase

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/constant"
	"gotu-bookstore/internal/contexts"
	"gotu-bookstore/internal/entity"
	"gotu-bookstore/mocks"
	"gotu-bookstore/pkg/authentication"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func TestUserUC_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyID := uuid.New()

	var dummyLongPassword []byte
	for i := 0; i < 80; i++ {
		dummyLongPassword = append(dummyLongPassword, byte(i))
	}

	type fields struct {
		userRepo internal.UserRepo
	}
	type args struct {
		req entity.UserRegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					mock.EXPECT().Store(gomock.Any(), gomock.Any()).Return(dummyID, nil)
					return mock
				}(),
			},
			args: args{
				req: entity.UserRegisterRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: nil,
		},
		{
			name: "error - password too long",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					return mock
				}(),
			},
			args: args{
				req: entity.UserRegisterRequest{
					Email:    "test@example.com",
					Password: string(dummyLongPassword),
				},
			},
			wantErr: bcrypt.ErrPasswordTooLong,
		},
		{
			name: "error - failed to store user",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					mock.EXPECT().Store(gomock.Any(), gomock.Any()).Return(uuid.Nil, errors.New("some error"))
					return mock
				}(),
			},
			args: args{
				req: entity.UserRegisterRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUC{
				logger:   zap.NewNop(),
				userRepo: tt.fields.userRepo,
			}
			err := u.Register(context.Background(), tt.args.req)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUserUC_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyID := uuid.New()

	type fields struct {
		cfg      *config.Cfg
		userRepo internal.UserRepo
		authJWT  authentication.AuthJWT
	}
	type args struct {
		req entity.AuthLoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				cfg: &config.Cfg{
					Auth: config.Auth{JwtSecret: "test-jwt-secret"},
				},
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{Email: "test@example.com"},
					}
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
					users := []entity.User{
						{
							BaseModel: entity.BaseModel{ID: dummyID},
							Password:  string(hashedPassword),
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(users, nil)
					return mock
				}(),
				authJWT: func() *mocks.MockAuthJWT {
					mock := mocks.NewMockAuthJWT(ctrl)
					mock.EXPECT().GenerateToken("test-jwt-secret", dummyID).Return("token-test-1234", nil)
					return mock
				}(),
			},
			args: args{
				req: entity.AuthLoginRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want:    "token-test-1234",
			wantErr: nil,
		},
		{
			name: "error - user not found",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{Email: "test@example.com"},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, constant.ErrNotFound)
					return mock
				}(),
			},
			args: args{
				req: entity.AuthLoginRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want:    "",
			wantErr: constant.ErrUserNotFound,
		},
		{
			name: "error - from user repo",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{Email: "test@example.com"},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, errors.New("some error"))
					return mock
				}(),
			},
			args: args{
				req: entity.AuthLoginRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want:    "",
			wantErr: errors.New("some error"),
		},
		{
			name: "error - invalid password",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{Email: "test@example.com"},
					}
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("different-password"), bcrypt.DefaultCost)
					users := []entity.User{
						{
							BaseModel: entity.BaseModel{ID: dummyID},
							Password:  string(hashedPassword),
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(users, nil)
					return mock
				}(),
			},
			args: args{
				req: entity.AuthLoginRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want:    "",
			wantErr: constant.ErrInvalidPassword,
		},
		{
			name: "error - from generate token",
			fields: fields{
				cfg: &config.Cfg{
					Auth: config.Auth{JwtSecret: "test-jwt-secret"},
				},
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{Email: "test@example.com"},
					}
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
					users := []entity.User{
						{
							BaseModel: entity.BaseModel{ID: dummyID},
							Password:  string(hashedPassword),
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(users, nil)
					return mock
				}(),
				authJWT: func() *mocks.MockAuthJWT {
					mock := mocks.NewMockAuthJWT(ctrl)
					mock.EXPECT().GenerateToken(gomock.Any(), dummyID).Return("", errors.New("some error from GenerateToken"))
					return mock
				}(),
			},
			args: args{
				req: entity.AuthLoginRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want:    "",
			wantErr: errors.New("some error from GenerateToken"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUC{
				cfg:      tt.fields.cfg,
				logger:   zap.NewNop(),
				userRepo: tt.fields.userRepo,
				authJWT:  tt.fields.authJWT,
			}
			got, err := u.Login(context.Background(), tt.args.req)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserUC_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyID := uuid.New()

	type fields struct {
		userRepo internal.UserRepo
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.UserResponse
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{
							BaseModel: entity.BaseModel{ID: dummyID},
						},
					}
					user := []entity.User{
						{
							BaseModel: entity.BaseModel{ID: dummyID},
							Email:     "test@example.com",
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(user, nil)
					return mock
				}(),
			},
			args: args{
				id: dummyID,
			},
			want: entity.UserResponse{
				ID:    dummyID.String(),
				Email: "test@example.com",
			},
			wantErr: nil,
		},
		{
			name: "error - not found",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{
							BaseModel: entity.BaseModel{ID: dummyID},
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, constant.ErrNotFound)
					return mock
				}(),
			},
			args: args{
				id: dummyID,
			},
			want:    entity.UserResponse{},
			wantErr: constant.ErrNotFound,
		},
		{
			name: "error - other error from user repo get",
			fields: fields{
				userRepo: func() *mocks.MockUserRepo {
					mock := mocks.NewMockUserRepo(ctrl)
					filter := entity.UserFilter{
						User: entity.User{
							BaseModel: entity.BaseModel{ID: dummyID},
						},
					}
					mock.EXPECT().Get(gomock.Any(), filter).Return(nil, errors.New("some error from user repo"))
					return mock
				}(),
			},
			args: args{
				id: dummyID,
			},
			want:    entity.UserResponse{},
			wantErr: errors.New("some error from user repo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUC{
				logger:   zap.NewNop(),
				userRepo: tt.fields.userRepo,
			}
			got, err := u.Get(context.Background(), tt.args.id)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserUC_Me(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyID := uuid.New()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    entity.UserResponse
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: func() context.Context {
					user := entity.User{
						BaseModel: entity.BaseModel{ID: dummyID},
					}
					ctx := context.WithValue(context.Background(), contexts.CtxKeyUser, user)
					return ctx
				}(),
			},
			want: entity.UserResponse{
				ID: dummyID.String(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUC{}
			got, err := u.Me(tt.args.ctx)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
