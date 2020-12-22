package routes

import (
	"goweibo/core"
	"goweibo/core/context"
	"goweibo/core/errno"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func registerError(app *core.Application) {
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler

	// 统一错误处理
	app.Engine.HTTPErrorHandler = func(err error, c echo.Context) {
		errnoData := transformErrorType(err)

		// Send response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(http.StatusOK)
			} else {
				// 响应错误的处理
				cc := context.NewAppContext(c)
				err = cc.ErrorResp(errnoData)
			}
			if err != nil {
				log.Printf("routes/error#HTTPErrorHandler: %s", err)
			}
		}
	}
}

func transformErrorType(err error) *errno.Errno {
	switch typed := err.(type) {
	// 请求参数错误
	case *errno.Errno:
		return typed
		// 其他 error
	default:
		return errno.UnknownErr.WithErr(typed).(*errno.Errno)
	}
}

func notFoundHandler(c echo.Context) error {
	return errno.NotFoundErr
}
