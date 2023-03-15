package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"UserService" summary:"Login with mobile and verify code"`
	Mobile string `v:"required|phone"`
	Code   string `v:"required|size:6"`
}
type UserLoginRes struct {
	Token   string `json:"token" dc:"the token with JWT format"`
	Expired int64  `json:"expired" dc:"expired time of this token"`
}

type UserJwtReq struct {
	g.Meta `path:"/jwt" method:"get"`
}
type UserJwtRes struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}

type UserProfileReq struct {
	g.Meta `path:"/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}
type UserProfileRes struct {
	*entity.User
}
