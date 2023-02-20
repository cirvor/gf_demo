package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
}

// ContextUser test
type ContextUser struct {
	UserId   uint   // User ID.
	Passport string // User passport.
	Nickname string // User nickname.
}
