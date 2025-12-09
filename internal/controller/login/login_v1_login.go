// internal/controller/login/login_v1_login.go
package login

import (
	"context"
	"leke/api/login"
	v1 "leke/api/login/v1"
	"leke/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ControllerV1 struct{}

func NewV1() login.ILoginV1 {
	return &ControllerV1{}
}

// Login 登录接口
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res, err = service.Login().Login(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{
		Token:   res.Token,
		Account: res.Account,
	}
	return res, nil
}

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	_, err = service.Login().Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterRes{
		Account: req.Account,
	}, nil
}

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

// 通过邮箱注册
func (c *ControllerV1) RegisterByEmail(ctx context.Context, req *v1.RegisterByEmailReq) (res *v1.RegisterByEmailRes, err error) {
	return service.Login().RegisterByEmail(ctx, req)
}

// 通过邮箱登录
func (c *ControllerV1) LoginByEmail(ctx context.Context, req *v1.LoginByEmailReq) (res *v1.LoginByEmailRes, err error) {
	return service.Login().LoginByEmail(ctx, req)
}

// 发送验证码
func (c *ControllerV1) SendVerificationCode(ctx context.Context, req *v1.SendVerificationCodeReq) (res *v1.SendVerificationCodeRes, err error) {
	return service.Login().SendVerificationCode(ctx, req)
}
