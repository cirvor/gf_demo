package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser test
type ContextUser struct {
	UserId   uint   // User ID.
	Nickname string // User nickname.
}
