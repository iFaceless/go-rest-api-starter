package controller

import (
	"database/sql"

	"github.com/ifaceless/go-starter/pkg/model"
	"github.com/pkg/errors"
)

func GetCompanies() ([]model.CompanyModel, error) {
	var companies []model.CompanyModel

	err := model.DB.Order("id").Find(&companies).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}

	return companies, nil
}
