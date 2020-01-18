package schema

import (
	"github.com/ifaceless/go-starter/pkg/admin/validator"
	"github.com/ifaceless/go-starter/pkg/model"
	"github.com/ifaceless/go-starter/pkg/util/pic"
	"github.com/ifaceless/go-starter/pkg/util/rest"
	"github.com/ifaceless/portal/field"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type OutputProductSchema struct {
	ID            *string               `json:"id,omitempty"`
	Title         *string               `json:"title,omitempty"`
	Intro         *string               `json:"intro,omitempty"`
	Description   *OutputDocumentSchema `json:"description,omitempty" portal:"nested;async"`
	IsInteractive *bool                 `json:"is_interactive,omitempty"`
	Artworks      *pic.Pictures         `json:"artwork_urls,omitempty" portal:"AUTO_INIT"`
	IsPublic      *bool                 `json:"is_public,omitempty"`
	PublicAt      *field.Timestamp      `json:"public_at,omitempty"`
	Company       *OutputCompanySchema  `json:"company,omitempty" portal:"nested;async"`
	CreatedAt     *field.Timestamp      `json:"created_at,omitempty"`
	UpdatedAt     *field.Timestamp      `json:"updated_at,omitempty"`
}

type InputProductSchema struct {
	Title         string              `json:"title" validate:"required"`
	Intro         string              `json:"intro" validate:"required,max=2048"`
	Description   InputDocumentSchema `json:"description" validate:"dive"`
	IsInteractive bool                `json:"is_interactive"`
	Artworks      pic.Pictures        `json:"artwork_urls"`
	Company       struct {
		ID string `json:"id"`
	} `json:"company"`
}

func (i *InputProductSchema) CompanyID() int64 {
	return cast.ToInt64(i.Company.ID)
}

func (i *InputProductSchema) Validate() error {
	err := i.ValidateCompany()
	if err != nil {
		return err
	}

	return validator.Validate.Struct(i)
}

func (i *InputProductSchema) ValidateCompany() error {
	var company model.CompanyModel
	id := cast.ToInt64(i.Company.ID)
	err := model.DB.Where("id = ?", id).Find(&company).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return rest.ResourceNotFound(errors.Errorf("公司 %s 不存在", i.Company.ID))
		}
		return err
	}
	return nil
}
