package handler

import (
	"github.com/ifaceless/go-starter/pkg/controller"
	"github.com/ifaceless/go-starter/pkg/util/rest"
	"github.com/ifaceless/go-starter/pkg/web/schema"
	"github.com/ifaceless/portal"
)

func GetCompanies(c *rest.Context) (rest.Response, error) {
	var outputs []schema.OutputCompanySchema

	companies, err := controller.GetCompanies()
	if err != nil {
		return nil, rest.ResourceNotFound(err)
	}

	err = portal.Dump(&outputs, companies)
	if err != nil {
		return nil, rest.BadRequest(err)
	}

	return outputs, nil
}
