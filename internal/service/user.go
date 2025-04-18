package service

import (
	"context"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

// UserService 是用户服务的实现
type UserService struct {
	v1.UnimplementedUserServiceServer

	uc *biz.UserUsecase
}

// NewUserService 创建一个新的用户服务实例
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// CreateUser 创建一个新的用户
func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	user := &biz.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	createdUser, err := s.uc.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &v1.CreateUserResponse{
		Id:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Age:       createdUser.Age,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(ctx context.Context, req *v1.GetUserByIDRequest) (*v1.GetUserByIDResponse, error) {
	user, err := s.uc.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserByIDResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	user := &biz.User{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	updatedUser, err := s.uc.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateUserResponse{
		Id:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Age:       updatedUser.Age,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	err := s.uc.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteUserResponse{}, nil
}

// ListUsers 列出所有用户
func (s *UserService) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	users, err := s.uc.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var pbUsers []*v1.User
	for _, user := range users {
		pbUsers = append(pbUsers, &v1.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Age:       user.Age,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return &v1.ListUsersResponse{Users: pbUsers}, nil
}