package user

import (
	"context"
	v1 "gf_demo/api/v1"
	"gf_demo/internal/model"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

// NLI is the controller for user.
var NLI = cUserNLI{}

type cUserNLI struct{}

// SignUp is the API for user sign up.
func (c *cUserNLI) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}

// SignIn is the API for user sign in.
func (c *cUserNLI) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}

// LoginIn
//
//	@Description: 用户通过手机号与验证码 登陆接口
//	@receiver c
//	@param ctx
//	@param req
//	@return res
//	@return err
func (c *cUserNLI) LoginIn(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	isExist, err := service.User().AuthMobileAndCode(ctx, &model.UserLoginInput{
		Mobile: req.Mobile,
		Code:   req.Code,
	})
	if err != nil {
		return
	}
	if !isExist {
		//err = gerror.NewCode(gcode.New(10000, "11111", nil), "用户校验失败")
		err = gerror.New("用户校验失败")
		return
	}
	return
}
