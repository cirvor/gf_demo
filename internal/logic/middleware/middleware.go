package middleware

import (
	"gf_demo/internal/model"
	"gf_demo/internal/service"
	"net/http"

	"github.com/gogf/gf/v2/util/gvalid"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// Auth
//
//	@Description: 调用jwt验证中间件
//	@receiver s
//	@param r
func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

// CORS
//
//	@Description: 跨域处理
//	@receiver s
//	@param r
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// ResponseHandler
//
//	@Description: 全局数据处理 + 异常捕获
//	@receiver s
//	@param r
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		// Validation error.
		if v, ok := err.(gvalid.Error); ok {
			// 处理
			r.Response.WriteJsonExit(model.ResponseData{
				Code:    66,
				Message: v.FirstError().Error(),
				Data:    res,
			})
		}

		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	r.Response.WriteJson(model.ResponseData{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
