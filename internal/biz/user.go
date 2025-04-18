package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)


// User 是用户模型
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int32     `json:"age"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

// UserRepo 定义了用户存储库的接口
type UserRepo interface {
	Create(context.Context, *User) (*User, error)
	GetByID(context.Context, int64) (*User, error)
	Update(context.Context, *User) (*User, error)
	Delete(context.Context, int64) error
	List(context.Context) ([]*User, error)
}

// UserUsecase 是用户业务逻辑的实现
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase 创建一个新的用户业务逻辑实例
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser 创建一个新的用户
func (uc *UserUsecase) CreateUser(ctx context.Context, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("Creating user: %v", user.Name)
	return uc.repo.Create(ctx, user)
}

// GetUserByID 根据ID获取用户
func (uc *UserUsecase) GetUserByID(ctx context.Context, id int64) (*User, error) {
	uc.log.WithContext(ctx).Infof("Getting user by ID: %d", id)
	return uc.repo.GetByID(ctx, id)
}

// UpdateUser 更新用户信息
func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("Updating user: %v", user.Name)
	return uc.repo.Update(ctx, user)
}

// DeleteUser 删除用户
func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("Deleting user by ID: %d", id)
	return uc.repo.Delete(ctx, id)
}

// ListUsers 列出所有用户
func (uc *UserUsecase) ListUsers(ctx context.Context) ([]*User, error) {
	uc.log.WithContext(ctx).Info("Listing all users")
	return uc.repo.List(ctx)
}
