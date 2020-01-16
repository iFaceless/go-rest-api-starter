package main

import (
	"encoding/json"
	"fmt"

	"github.com/ifaceless/go-starter/pkg/util/seqgen"

	"github.com/ifaceless/go-starter/pkg/model"
	"github.com/ifaceless/portal"
	"github.com/joho/godotenv"

	"github.com/ifaceless/go-starter/pkg/admin/schema"
)

var companies = `
[
    {
        "title": "知乎",
        "intro": "可信赖的问答社区，以让每个人高效获得可信赖的解答为使命。",
        "artwork_urls": [
            "https://pic2.zhimg.com/v2-f8f2d20ee4cd5a1cc87c4ddba58f0078.jpg"
        ],
        "url": "https://www.zhihu.com"
    },
    {
        "title": "AIZOO",
        "intro": "AI 算法商城，还有算法乐园等你体验",
        "artwork_urls": [
            "https://pic2.zhimg.com/v2-f8f2d20ee4cd5a1cc87c4ddba581223444.jpg"
        ],
        "url": "https://aizoo.com"
    }
]
`

func main() {
	_ = godotenv.Load(".env")
	var companySchemas []schema.InputCompanySchema
	err := json.Unmarshal([]byte(companies), &companySchemas)
	if err != nil {
		panic(err)
	}
	fmt.Println(companySchemas)

	for _, cs := range companySchemas {
		var companyModel model.CompanyModel
		// Portal 会自动完成 schema 和 model 的映射
		err := portal.Dump(&companyModel, cs)
		if err != nil {
			panic(err)
		}

		id, _ := seqgen.NextID()
		companyModel.ID = id

		fmt.Println(companyModel)
		model.DB.Create(&companyModel)
	}
}
