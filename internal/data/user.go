package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "hello/api/user/v1"
	"hello/internal/biz"
	"hello/internal/data/entity"
	"hello/internal/pkg/timehelper"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "moudle", "repo/user")),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	entity := entity.User{
		Name:   user.Name,
		Age:    user.Age,
		Mobile: user.Mobile,
	}

	if err := r.data.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return nil, err
	}

	return &biz.User{
		Id:     entity.Id,
		Name:   entity.Name,
		Age:    entity.Age,
		Mobile: entity.Mobile,
	}, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	var entity entity.User
	if err := r.data.db.WithContext(ctx).First(&entity, user.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorUserNotFound("用户 %d 没找到", user.Id)
		}
		return nil, err
	}

	entity.Name = user.Name
	entity.Age = user.Age
	entity.Mobile = user.Mobile

	if err := r.data.db.WithContext(ctx).Save(&entity).Error; err != nil {
		return nil, err
	}

	return entityToDemoDomain(&entity), nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int64) error {
	var entity entity.User
	if err := r.data.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return v1.ErrorUserNotFound("用户 %d 没找到", id)
		}
		return err
	}

	return r.data.db.WithContext(ctx).Delete(&entity).Error
}

func (r *UserRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	var repoDemo entity.User
	if err := r.data.db.WithContext(ctx).First(&repoDemo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorUserNotFound("用户 %d 没找到", id)
		}
		return nil, err
	}

	return entityToDemoDomain(&repoDemo), nil
}

func (r *UserRepo) ListUser(ctx context.Context, req *biz.ListUser) ([]*biz.User, int64, error) {
	conn := r.data.db.WithContext(ctx).Model(&entity.User{})

	if req.Mobile != "" {
		conn = conn.Where("mobile = ?", req.Mobile)
	}

	if req.CreatedStart != "" {
		conn = conn.Where("created_at >= ?", req.CreatedStart)
	}

	if req.CreatedEnd != "" {
		conn = conn.Where("created_at <= ?", req.CreatedEnd)
	}

	var count int64
	conn.Count(&count)

	var repoDemos []*entity.User
	if err := conn.Scopes(entity.Paginate(req.Page, req.PageSize)).Find(&repoDemos).Error; err != nil {
		return nil, 0, err
	}

	var responseList []*biz.User
	for _, demo := range repoDemos {
		responseList = append(responseList, entityToDemoDomain(demo))
	}

	return responseList, count, nil
}

func entityToDemoDomain(entity *entity.User) *biz.User {
	return &biz.User{
		Id:        entity.Id,
		Name:      entity.Name,
		Age:       entity.Age,
		Mobile:    entity.Mobile,
		CreatedAt: timehelper.FormatYMDHIS(&entity.CreatedAt),
		UpdatedAt: timehelper.FormatYMDHIS(&entity.UpdatedAt),
	}
}
