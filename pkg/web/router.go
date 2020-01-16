package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ifaceless/go-starter/pkg/middleware"
	"github.com/ifaceless/go-starter/pkg/util/rest"
	"github.com/ifaceless/go-starter/pkg/web/handler"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.DefaultCORS())

	v1 := r.Group("/v1")
	{
		v1.GET("/companies", rest.Adapt(handler.GetCompanies))
	}
	return r
}
