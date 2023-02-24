package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RedisTestReq struct {
	g.Meta `path:"/redis-test" tags:"test" method:"get" summary:"You first hello api"`
}
type RedisTestRes struct{}
