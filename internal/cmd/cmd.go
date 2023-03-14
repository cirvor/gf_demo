package cmd

import (
	"context"
	"gf_demo/internal/consts"
	"gf_demo/internal/controller"
	"gf_demo/internal/controller/goods"
	"gf_demo/internal/controller/user"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/net/goai"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 设置全局中间件
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)

				group.Bind(
					controller.Redis,
				)

				// 路由注册用户模块
				group.Group("/user", func(group *ghttp.RouterGroup) {
					// 注册用户模块
					group.Bind(
						user.NLI,
					)
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(service.Middleware().Auth)
						group.Bind(
							user.User,
						)
					})
				})

				// 路由注册商品模块
				group.Group("/goods", func(group *ghttp.RouterGroup) {
					// 注册用户模块
					group.Bind(
						goods.NLI,
					)
					//group.Group("/", func(group *ghttp.RouterGroup) {
					//	group.Middleware(service.Middleware().Auth)
					//	group.Bind(
					//		user.User,
					//	)
					//})
				})

			})

			// 额外处理接口文档显示数据
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "GoFrame",
			URL:  "https://goframe.org",
		},
	}
}
