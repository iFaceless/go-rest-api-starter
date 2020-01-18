package handler

import (
	"github.com/ifaceless/go-starter/pkg/controller"
	"github.com/ifaceless/go-starter/pkg/util/rest"
	"github.com/ifaceless/go-starter/pkg/web/schema"
	"github.com/ifaceless/portal"
)

func GetCompanies(_ *rest.Context) (rest.Response, error) {
	var outputs []schema.OutputCompanySchema

	// 调用 controller 层接口获得结果
	companies, err := controller.GetCompanies()
	if err != nil {
		return nil, rest.ResourceNotFound(err)
	}

	// 借助 portal 将数据序列化为指定的 Schema 结构体
	err = portal.Dump(&outputs, companies)
	if err != nil {
		return nil, rest.BadRequest(err)
	}

	// 直接返回结果，rest 框架会自动使用 json 序列化
	return outputs, nil
}
