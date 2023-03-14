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

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			UserId:   user.UserId,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
}

// Auth validates the request to allow only signed-in users visit.
func (s *sMiddleware) Auth(r *ghttp.Request) {
	if service.User().IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// HandlerResponse is the default middleware handling handler response object and its error.
func (s *sMiddleware) HandlerResponse(r *ghttp.Request) {
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
			code = gcode.CodeInvalidRequest
			msg = v.FirstError().Error()
		} else {
			if code == gcode.CodeNil {
				code = gcode.CodeInternalError
			}
			msg = err.Error()
		}
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
