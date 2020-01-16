// Package constant 定义一些常量或枚举
//go:generate enum -type=CategoryGroupKind -linecomment=true -output=gen_enum.go
package constant

type CategoryGroupKind int

const (
	CategoryProductType  CategoryGroupKind = iota + 1 // product_type
	CategoryTechType                                  // tech_type
	CategoryIndustryType                              // industry_type
)

func (i CategoryGroupKind) Title() string {
	switch i {
	case CategoryProductType:
		return "产品类型"
	case CategoryTechType:
		return "技术类型"
	case CategoryIndustryType:
		return "应用行业"
	default:
		return "未知分组"
	}
}
