package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "hello/api/user/v1"
	"hello/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer

	userUc *biz.UserUseCase
	log    *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		userUc: uc,
		log:    log.NewHelper(log.With(logger, "module", "service/user")),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user := &biz.User{
		Name:   req.User.Name,
		Age:    req.User.Age,
		Mobile: req.User.Mobile,
	}

	user, err := s.userUc.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{Id: user.Id}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	user := &biz.User{
		Id:     req.User.Id,
		Name:   req.User.Name,
		Age:    req.User.Age,
		Mobile: req.User.Mobile,
	}
	if _, err := s.userUc.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := s.userUc.DeleteUser(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.userUc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{User: &pb.UserEntity{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
		Mobile:    user.Mobile,
		CreatedAt: user.CreatedAt,
	}}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	reqList := &biz.ListUser{
		Page:         req.Page,
		PageSize:     req.PageSize,
		Mobile:       req.Mobile,
		CreatedStart: req.CreatedStart,
		CreatedEnd:   req.CreatedEnd,
	}

	var userList []*pb.UserEntity
	userDomainList, count, err := s.userUc.ListUser(ctx, reqList)
	if err != nil {
		return nil, err
	}

	for _, demo := range userDomainList {
		userList = append(userList, userDomainToReply(demo))
	}

	return &pb.ListUserReply{Total: count, Users: userList}, nil
}

func userDomainToReply(user *biz.User) *pb.UserEntity {
	return &pb.UserEntity{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
		Mobile:    user.Mobile,
		CreatedAt: user.CreatedAt,
	}
}
