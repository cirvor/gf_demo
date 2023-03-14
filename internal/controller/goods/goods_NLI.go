package goods

import (
	"context"
	v1 "gf_demo/api/v1"
	"gf_demo/internal/model"
	"gf_demo/internal/service"
)

// NLI is the controller for user.
var NLI = cGoodsNLI{}

type cGoodsNLI struct{}

// Add
//
//	@Description: 添加商品接口
//	@receiver c
//	@param ctx
//	@param req
//	@return res
//	@return err
func (c *cGoodsNLI) Add(ctx context.Context, req *v1.GoodsAddReq) (res *v1.GoodsAddRes, err error) {
	err = service.Goods().Add(ctx, &model.GoodsAddInput{
		Name:        req.Name,
		Description: req.Description,
	})
	return
}

// Info
//
//	@Description: 获取商品信息接口
//	@receiver c
//	@param ctx
//	@param req
//	@return res
//	@return err
func (c *cGoodsNLI) Info(ctx context.Context, req *v1.GoodsInfoReq) (res *v1.GoodsInfoRes, err error) {
	// 查找产品信息
	res = &v1.GoodsInfoRes{
		Goods: service.Goods().Info(ctx, &model.GoodsIdInInput{Id: req.Id}),
	}

	return
}
