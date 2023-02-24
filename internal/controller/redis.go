package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/v1"
)

var (
	Redis = cRedis{}
)

type cRedis struct{}

func (c *cRedis) Test(ctx context.Context, req *v1.RedisTestReq) (res *v1.RedisTestRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
