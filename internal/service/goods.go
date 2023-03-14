// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_demo/internal/model"
	"gf_demo/internal/model/entity"
)

type (
	IGoods interface {
		Info(ctx context.Context, in *model.GoodsIdInInput) (goods *entity.Goods)
		Add(ctx context.Context, in *model.GoodsAddInput) (err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
