package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"hello/internal/pkg/validate"
)

type User struct {
	Id        int64
	Name      string `validate:"required,len=10" label:"姓名"`
	Age       int32
	Mobile    string `validate:"required"`
	CreatedAt string
	UpdatedAt string
}

type ListUser struct {
	Page         int64
	PageSize     int64
	Name         string
	Mobile       string
	CreatedStart string
	CreatedEnd   string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*User, error)
	ListUser(ctx context.Context, list *ListUser) ([]*User, int64, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	if err := validate.Validate(user); err != nil {
		return nil, err
	}
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) ListUser(ctx context.Context, list *ListUser) ([]*User, int64, error) {
	return uc.repo.ListUser(ctx, list)
}
