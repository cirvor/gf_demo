package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GoodsAddReq struct {
	g.Meta      `path:"/add" tags:"GoodsService" method:"post" dc:"add goods"`
	Name        string `v:"required|length:3,16"`
	Description string `v:"required|length:6,16"`
}

type GoodsAddRes struct{}

type GoodsInfoReq struct {
	g.Meta `path:"/info" tags:"GoodsService" method:"get" dc:"show goods info"`
	Id     int `v:"required"`
}

type GoodsInfoRes struct {
	*entity.Goods
}
