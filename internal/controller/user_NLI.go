package controller

import (
	"context"
	v1 "gf_demo/api/v1"
	"gf_demo/internal/model"
	"gf_demo/internal/service"
)

// UserNLI is the controller for user.
var UserNLI = cUserNLI{}

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
