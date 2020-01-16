package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ifaceless/go-starter/pkg/admin/handler"
	"github.com/ifaceless/go-starter/pkg/middleware"
	"github.com/ifaceless/go-starter/pkg/util/rest"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.DefaultCORS())

	v1 := r.Group("/v1")
	{
		productAPI := v1.Group("/products")
		{
			productAPI.GET("", rest.Adapt(handler.GetProducts))
			productAPI.GET("/:product_id", rest.Adapt(handler.GetProduct))
			productAPI.POST("", rest.Adapt(handler.CreateProduct))
			productAPI.PUT("/:product_id", rest.Adapt(handler.UpdateProduct))
			productAPI.DELETE("/:product_id", rest.Adapt(handler.DeleteProduct))
			productAPI.POST("/:product_id/public", rest.Adapt(handler.PublicProduct))
		}
	}
	return r
}
