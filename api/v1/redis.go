package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RedisTestReq struct {
	g.Meta `path:"/redis-test" tags:"test" method:"get" dc:"test redis"`
}

type RedisTestRes struct {
	Time string `json:"time" dc:"get the val in redis"`
}
