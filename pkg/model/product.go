package model

import (
	"fmt"
	"time"

	"github.com/ifaceless/go-starter/pkg/model/column"
	"github.com/ifaceless/go-starter/pkg/util/pic"
	"github.com/jinzhu/gorm"

	"github.com/pkg/errors"
)

type ProductModel struct {
	ID        int64
	CompanyID int64
	Title     string
	Intro     string
	DescDocID int64
	// Artworks only save picture token to db, but return with full urls for API.
	Artworks      *pic.Pictures
	IsInteractive column.Bool
	IsPublic      column.Bool
	PublicAt      *time.Time
	Extra         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (pd *ProductModel) TableName() string {
	return "product"
}

func (pd *ProductModel) String() string {
	return fmt.Sprintf("<ProductModel id=%d>", pd.ID)
}

func (pd *ProductModel) Self() *ProductModel {
	return pd
}

func (pd *ProductModel) Description() (*DocumentModel, error) {
	var doc DocumentModel
	result := DB.Where("id = ?", pd.DescDocID).First(&doc)
	if result.Error != nil && !gorm.IsRecordNotFoundError(result.Error) {
		return nil, errors.WithStack(result.Error)
	}

	return &doc, nil
}

func (pd *ProductModel) Company() (*CompanyModel, error) {
	var company CompanyModel
	err := DB.Where("id = ?", pd.CompanyID).Find(&company).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &company, nil
}
