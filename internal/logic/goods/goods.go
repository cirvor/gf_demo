package goods

import (
	"context"
	"gf_demo/internal/dao"
	"gf_demo/internal/model"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
)

type (
	sGoods struct{}
)

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

func (s *sGoods) Info(ctx context.Context, in *model.GoodsIdInInput) (goods *entity.Goods) {
	err := dao.Goods.Ctx(ctx).WherePri(in.Id).Scan(&goods)
	if err != nil {
		return nil
	}
	return
}

func (s *sGoods) Add(ctx context.Context, in *model.GoodsAddInput) (err error) {
	_, err = dao.Goods.Ctx(ctx).Data(do.Goods{
		Name:        in.Name,
		Description: in.Description,
	}).Insert()
	return err
}
