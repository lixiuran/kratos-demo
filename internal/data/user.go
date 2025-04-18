package data

import (
	"context"
	"database/sql"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo 是 UserRepo 接口的具体实现
type UserRepo struct {
	db  *sql.DB
	log *log.Helper
}

// NewUserRepo 创建一个新的 UserRepo 实例
func NewUserRepo(db *sql.DB, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// Create 创建一个新的用户
func (r *UserRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Age)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}

// GetByID 根据ID获取用户
func (r *UserRepo) GetByID(ctx context.Context, id int64) (*biz.User, error) {
	query := "SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var user biz.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, biz.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// Update 更新用户信息
func (r *UserRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	query := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete 删除用户
func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// List 列出所有用户
func (r *UserRepo) List(ctx context.Context) ([]*biz.User, error) {
	query := "SELECT id, name, email, age, created_at, updated_at FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*biz.User
	for rows.Next() {
		var user biz.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}