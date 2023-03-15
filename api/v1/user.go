package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserProfileReq struct {
	g.Meta `path:"/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}
type UserProfileRes struct {
	*entity.User
}

type UserSignUpReq struct {
	g.Meta    `path:"/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}
type UserSignUpRes struct{}

type UserSignInReq struct {
	g.Meta   `path:"/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Passport string `v:"required"`
	Password string `v:"required"`
}
type UserSignInRes struct{}

type UserCheckPassportReq struct {
	g.Meta   `path:"/check-passport" method:"post" tags:"UserService" summary:"Check passport available"`
	Passport string `v:"required"`
}
type UserCheckPassportRes struct{}

type UserCheckNickNameReq struct {
	g.Meta   `path:"/check-nickname" method:"post" tags:"UserService" summary:"Check nickname available"`
	Nickname string `v:"required"`
}
type UserCheckNickNameRes struct{}

type UserIsSignedInReq struct {
	g.Meta `path:"/is-signed-in" method:"post" tags:"UserService" summary:"Check current user is already signed-in"`
}
type UserIsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type UserSignOutReq struct {
	g.Meta `path:"/sign-out" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type UserSignOutRes struct{}

type UserLoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"UserService" summary:"Login with mobile and verify code"`
	Mobile string `v:"required|phone"`
	Code   string `v:"required|size:6"`
}
type UserLoginRes struct {
	Token   string `json:"token" dc:"the token with JWT format"`
	Expired int64  `json:"expired" dc:"expired time of this token"`
}
