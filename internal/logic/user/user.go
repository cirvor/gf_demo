package user

import (
	"context"
	"gf_demo/internal/dao"
	"gf_demo/internal/model"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// IsMobileExist
//
//	@Description: 校验当前手机号是否存在
//	@receiver s
//	@param ctx
//	@param mobile
//	@return bool
//	@return error
func (s *sUser) IsMobileExist(ctx context.Context, mobile string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Mobile: mobile,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sUser) GetProfile(ctx context.Context) (user *entity.User, err error) {
	userId := service.Auth().GetIdentity(ctx)
	//glog.Warning(ctx, userId)

	err = dao.User.Ctx(ctx).WherePri(userId).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		err = gerror.New("用户数据异常！请联系客服")
	}

	return
}

// AuthMobileAndCode
//
//	@Description: 通过手机号与验证码验证登陆信息
//	@receiver s
//	@param ctx
//	@return *entity.User
func (s *sUser) AuthMobileAndCode(ctx context.Context, in *model.UserLoginInput) (bool, error) {
	// todo 校验验证码
	if in.Code != "123456" {
		return false, nil
	}

	// todo 改为find查找
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Mobile: in.Mobile,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// GetUserByMobile
//
//	@Description: 通过手机号查询用户信息
//	@receiver s
//	@param ctx
//	@param in
//	@return map[string]interface{}
func (s *sUser) GetUserByMobile(ctx context.Context, in *model.UserLoginInput) map[string]interface{} {
	user := entity.User{}
	err := dao.User.Ctx(ctx).Where(do.User{
		Mobile: in.Mobile,
	}).Scan(&user)
	if err != nil {
		return nil
	}

	return g.Map{
		"id":       user.UserId,
		"username": user.Nickname,
	}
}
