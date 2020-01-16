package handler

import (
	"github.com/ifaceless/go-starter/pkg/controller"

	"github.com/gin-gonic/gin"
	"github.com/ifaceless/go-starter/pkg/admin/schema"
	"github.com/ifaceless/go-starter/pkg/model"
	"github.com/ifaceless/go-starter/pkg/util/rest"
	"github.com/ifaceless/portal"
	"github.com/spf13/cast"
)

func GetProducts(c *rest.Context) (rest.Response, error) {
	products, total, err := controller.GetProducts(
		c.Query("title"),
		c.QueryWithFallback("order_by", "-created_at"),
		c.Offset(),
		c.Limit(),
	)

	var schemas []schema.OutputProductSchema
	err = portal.Dump(&schemas, products)
	if err != nil {
		return nil, err
	}

	return rest.NewPage(c, schemas, total), nil
}

func CreateProduct(c *rest.Context) (rest.Response, error) {
	var input schema.InputProductSchema
	err := c.BindJSON(&input)
	if err != nil {
		return nil, err
	}

	var product model.ProductModel
	err = portal.Dump(&product, input)
	if err != nil {
		return nil, err
	}

	prodID, err := controller.CreateProduct(&product, &input)
	if err != nil {
		return nil, err
	}

	return gin.H{"success": true, "product_id": cast.ToString(prodID)}, nil
}

func GetProduct(c *rest.Context) (rest.Response, error) {
	product, err := controller.GetProduct(c.IntParam("product_id"))
	if err != nil {
		return nil, err
	}

	var output schema.OutputProductSchema
	err = portal.Dump(&output, product)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func UpdateProduct(c *rest.Context) (rest.Response, error) {
	product, err := controller.GetProduct(c.IntParam("product_id"))
	if err != nil {
		return nil, err
	}

	var input schema.InputProductSchema
	err = c.BindJSON(&input)
	if err != nil {
		return nil, err
	}

	err = portal.Dump(&product, input)
	if err != nil {
		return nil, err
	}

	err = controller.UpdateProduct(product, &input)
	if err != nil {
		return nil, err
	}

	return gin.H{"success": true}, nil
}

func DeleteProduct(c *rest.Context) (rest.Response, error) {
	product, err := controller.GetProduct(c.IntParam("product_id"))
	if err != nil {
		return nil, err
	}

	err = controller.DeleteProduct(product)
	if err != nil {
		return nil, err
	}

	return gin.H{"success": true}, nil
}

func PublicProduct(c *rest.Context) (rest.Response, error) {
	prod, err := controller.GetProduct(c.IntParam("product_id"))
	if err != nil {
		return nil, err
	}

	input := struct {
		IsPublic bool `json:"is_public"`
	}{}

	err = c.BindJSON(&input)
	if err != nil {
		return nil, err
	}

	err = controller.PublicProduct(prod, input.IsPublic)
	if err != nil {
		return nil, err
	}
	return gin.H{"success": true}, nil
}
