package user

import (
	"context"
	v1 "gf_demo/api/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// User is the controller for user.
var User = cUser{}

type cUser struct{}

// Profile
//
//	@Description: 用户信息接口
//	@receiver c
//	@param ctx
//	@param req
//	@return res
//	@return err
func (c *cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	user, err := service.User().GetProfile(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UserProfileRes{
		User: user,
	}
	return
}

// JWT
//
//	@Description: 查看jwt数据
//	@receiver c
//	@param ctx
//	@param req
//	@return res
//	@return err
func (c *cUser) JWT(ctx context.Context, req *v1.UserJwtReq) (res *v1.UserJwtRes, err error) {
	return &v1.UserJwtRes{
		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
		IdentityKey: service.Auth().IdentityKey,
		Payload:     service.Auth().GetPayload(ctx),
	}, nil
}
