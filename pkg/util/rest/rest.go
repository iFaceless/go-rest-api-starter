package rest

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

type (
	Response   interface{}
	HandleFunc func(c *Context) (Response, error)
)

func Adapt(fn HandleFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				fmt.Println(string(buf[:n]))
				abortWithError(context, InternalServerError(errors.Errorf("%s", p)))
			}
		}()

		data, err := fn(&Context{context})
		if err != nil {
			abortWithError(context, err)
		} else {
			context.JSON(http.StatusOK, data)
		}
	}
}
