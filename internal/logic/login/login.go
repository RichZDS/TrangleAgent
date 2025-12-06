// Package login internal/logic/login/login.go
package login

import (
	"context"
	"errors"
	v1 "leke/api/login/v1"
	"leke/internal/consts"
	"leke/internal/dao"
	"leke/internal/middleware"
	"leke/internal/model"
	"leke/internal/service"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/golang-jwt/jwt/v5"
)

// 注意：名字改成 sLogin，首字母小写 s + 大写 Login
type sLogin struct{}

func New() *sLogin {
	return &sLogin{}
}

// 关键！！在 init 里完成注册
func init() {
	service.RegisterLogin(New())
}

func (s *sLogin) Register(ctx context.Context, loginReq *v1.RegisterReq) (res *v1.RegisterRes, err error) {

	// 判断参数是否正常 账户的长度要在5-10 密码要经过MD5加密
	if len(loginReq.Account) < 5 || len(loginReq.Account) > 10 {
		glog.Debugf(ctx, "账户长度校验失败，长度：%d", len(loginReq.Account))
		return nil, errors.New("账户长度要在5-10之间")
	}
	if len(loginReq.Password) < 6 || len(loginReq.Password) > 32 {
		glog.Debugf(ctx, "密码长度校验失败，长度：%d", len(loginReq.Password))
		return nil, errors.New("密码长度要在6-32之间")
	}

	password, err := gmd5.Encrypt(loginReq.Password + consts.Salt)
	if err != nil {
		glog.Errorf(ctx, "密码加密失败：%v", err)
		return nil, gerror.Wrap(err, "密码加密失败，请稍后重试！")
	}

	loginData := model.LoginField{
		Account:  loginReq.Account,
		Password: password,
		Nickname: "好好先生",
	}

	_, err = dao.Users.Ctx(ctx).Data(loginData).Insert()
	if err != nil {
		glog.Errorf(ctx, "插入账号到数据库失败：%v", err)
		return nil, gerror.Wrap(err, "插入账号到数，据库失败，请稍后重试！")
	}
	return
}

// JWTClaims represents the custom claims for the JWT token
type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *sLogin) Login(ctx context.Context, loginReq *v1.LoginReq) (res *v1.LoginRes, err error) {

	//校验参数
	if len(loginReq.Account) < 5 || len(loginReq.Account) > 10 {
		glog.Debugf(ctx, "账户长度校验失败，长度：%d", len(loginReq.Account))
		return nil, errors.New("账户长度要在5-10之间")
	}
	if len(loginReq.Password) < 6 || len(loginReq.Password) > 32 {
		glog.Debugf(ctx, "密码长度校验失败，长度：%d", len(loginReq.Password))
		return nil, errors.New("密码长度要在6-32之间")
	}

	// 2. 先根据账号查数据库
	var user model.LoginField
	err = dao.Users.Ctx(ctx).Where("account", loginReq.Account).Scan(&user)
	if err != nil {
		glog.Errorf(ctx, "查询用户失败：%v", err)
		return nil, gerror.Wrap(err, "登录失败，请稍后重试！")
	}
	if user.Account == "" {
		// 没查到用户
		return nil, gerror.New("账号或密码错误")
	}

	// 3. 使用“相同规则”加密用户输入的密码，并和 DB 中的密码对比
	// 注意：这里的加密规则必须和 Register 时一致
	encryptedInput, err := gmd5.Encrypt(loginReq.Password + consts.Salt)
	if err != nil {
		glog.Errorf(ctx, "密码加密失败：%v", err)
		return nil, gerror.Wrap(err, "密码加密失败，请稍后重试！")
	}
	if user.Password != encryptedInput {
		return nil, gerror.New("账号或密码错误")
	}

	// 5. 种 Session（如果你确实需要 session）
	if err := setSession(ctx, user.Account); err != nil {
		glog.Errorf(ctx, "设置Session失败：%v", err)
		return nil, gerror.Wrap(err, "设置登录状态失败，请稍后重试！")
	}
	//JWT 生成
	// Create claims with user information
	claims := &JWTClaims{
		Username: loginReq.Account,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 生成Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(middleware.JwtSecretKey))
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "Failed to generate token")
	}
	res = &v1.LoginRes{
		Token:   signedToken,
		Account: user.Account,
	}
	return
}

func (s *sLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	// 从上下文里拿到当前请求
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.New("无效的请求上下文")
	}

	// 1. 删除当前登录用到的 Session 信息
	//    和 Login / setSession 里保持一致：都是用 "userAccount" 这个 key
	if err = r.Session.Remove("userAccount"); err != nil {
		glog.Errorf(ctx, "删除Session失败：%v", err)
		return nil, gerror.Wrap(err, "退出登录失败，请稍后重试！")
	}

	// 如果你想彻底清空这个 Session（不只删一个字段），可以用：
	// if err = r.Session.RemoveAll(); err != nil { ... }

	// 2. 构造返回结果（一般是一个空结构就够了）
	res = &v1.LogoutRes{}
	return res, nil
}

// 给前端种session
func setSession(ctx context.Context, account string) error {
	return g.RequestFromCtx(ctx).Session.Set("userAccount", account)
}
