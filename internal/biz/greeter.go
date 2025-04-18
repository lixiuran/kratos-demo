package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterRepo 定义了问候服务的存储库接口
type GreeterRepo interface {
	SayHello(ctx context.Context, name string) (string, error)
}

// Greeter 是问候服务的实现
type Greeter struct {
	Hello string
	log   *log.Helper
}

// GreeterUsecase 是问候服务的用例
type GreeterUsecase struct {
	log *log.Helper
}

// NewGreeterUsecase 创建一个新的问候服务用例
func NewGreeterUsecase(logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{
		log: log.NewHelper(logger),
	}
}

// CreateGreeter 创建问候
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("Creating greeter: %v", g)
	return g, nil
}

// NewGreeter 创建一个新的问候服务实例
func NewGreeter(logger log.Logger) *Greeter {
	return &Greeter{log: log.NewHelper(logger)}
}

// SayHello 发送问候消息
func (g *Greeter) SayHello(ctx context.Context, name string) (string, error) {
	g.log.WithContext(ctx).Infof("Saying hello to: %s", name)
	return "Hello, " + name, nil
}