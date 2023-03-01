package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-module/carbon/v2"

	v1 "gf_demo/api/v1"
)

var (
	Redis = cRedis{}
)

type cRedis struct{}

func (c *cRedis) Test(ctx context.Context, req *v1.RedisTestReq) (res *v1.RedisTestRes, err error) {
	_, err = g.Redis().Set(ctx, "key", carbon.Now().ToDateTimeString())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	value, err := g.Redis().Get(ctx, "key")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	fmt.Println(value.String())

	res = &v1.RedisTestRes{
		Time: value.String(),
	}

	return
}
