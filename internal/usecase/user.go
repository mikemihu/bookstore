package usecase

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/constant"
	"gotu-bookstore/internal/contexts"
	"gotu-bookstore/internal/entity"
	authPkg "gotu-bookstore/pkg/authentication"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserUC struct {
	cfg      *config.Cfg
	logger   *zap.Logger
	userRepo internal.UserRepo
	authJWT  authPkg.AuthJWT
}

func NewUserUC(
	cfg *config.Cfg,
	logger *zap.Logger,
	userRepo internal.UserRepo,
	authJWT authPkg.AuthJWT,
) internal.UserUC {
	return &UserUC{
		cfg:      cfg,
		logger:   logger,
		userRepo: userRepo,
		authJWT:  authJWT,
	}
}

func (u *UserUC) Register(ctx context.Context, req entity.UserRegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	user.ID, err = u.userRepo.Store(ctx, user)
	if err != nil {
		u.logger.Error("failed to store user", zap.Error(err),
			zap.Any("user", user))
		return err
	}

	return nil
}

func (u *UserUC) Login(ctx context.Context, req entity.AuthLoginRequest) (string, error) {
	filter := entity.UserFilter{
		User: entity.User{Email: req.Email},
	}
	users, err := u.userRepo.Get(ctx, filter)
	if errors.Is(err, constant.ErrNotFound) {
		return "", constant.ErrUserNotFound
	}
	if err != nil {
		u.logger.Error("failed to get user", zap.Error(err),
			zap.Any("user", req))
		return "", err
	}
	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", constant.ErrInvalidPassword
	}

	token, err := u.authJWT.GenerateToken(user.ID)
	if err != nil {
		u.logger.Error("failed authJWT.GenerateToken", zap.Error(err),
			zap.Any("user", user))
		return "", err
	}

	return token, nil
}

func (u *UserUC) Get(ctx context.Context, id uuid.UUID) (entity.UserResponse, error) {
	filter := entity.UserFilter{
		User: entity.User{
			BaseModel: entity.BaseModel{ID: id},
		},
	}
	users, err := u.userRepo.Get(ctx, filter)
	if errors.Is(err, constant.ErrNotFound) {
		return entity.UserResponse{}, constant.ErrNotFound
	}
	if err != nil {
		u.logger.Error("failed to get user", zap.Error(err),
			zap.Any("filter", filter))
		return entity.UserResponse{}, err
	}
	return users[0].ToResponse(), nil
}

func (_ *UserUC) Me(ctx context.Context) (entity.UserResponse, error) {
	user := contexts.GetUser(ctx)
	return user.ToResponse(), nil
}
