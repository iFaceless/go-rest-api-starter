package controller

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ifaceless/go-starter/pkg/util/rest"

	"github.com/ifaceless/go-starter/pkg/util"

	"github.com/ifaceless/go-starter/pkg/util/seqgen"

	"github.com/ifaceless/go-starter/pkg/admin/schema"
	"github.com/ifaceless/go-starter/pkg/model"
	"github.com/ifaceless/portal"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	allowedOrderByColumns = []string{"id", "product_id", "created_at", "is_public", "public_at", "is_interactive"}
)

func GetProducts(title string, orderBy string, offset, limit int) ([]model.ProductModel, int, error) {
	var products []model.ProductModel
	var total int

	orderBy, err := util.ParseOrderBy(orderBy, allowedOrderByColumns)
	if err != nil {
		return nil, 0, err
	}

	db := model.DB.Order(orderBy)
	if title != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", title))
	}

	db.Model(&products).Count(&total)
	err = db.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil && err != sql.ErrNoRows {
		return nil, 0, err
	}

	return products, total, nil
}

func GetProduct(productID int) (*model.ProductModel, error) {
	var product model.ProductModel
	err := model.DB.Where("id = ?", productID).Find(&product).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, rest.ResourceNotFound(errors.Errorf("产品 %d 不存在", productID))
		}
		return nil, err
	}

	return &product, nil
}

func CreateProduct(product *model.ProductModel, input *schema.InputProductSchema) (int64, error) {
	prodID, err := seqgen.NextID()
	if err != nil {
		return 0, err
	}

	docID, err := seqgen.NextID()
	if err != nil {
		return 0, err
	}

	err = model.Transact(func(tx *gorm.DB) error {
		logrus.Infof("[controller.CreateProduct] generated product id %d, doc id %d", prodID, docID)

		// 创建产品
		product.ID = prodID
		product.DescDocID = docID
		logrus.Infof("[controller.CreateProduct] create product: %s", product)
		err = tx.Create(&product).Error
		if err != nil {
			return err
		}

		// 创建文档
		doc := model.DocumentModel{
			ID:              docID,
			ContentHTML:     input.Description.ContentHTML,
			ContentMarkdown: input.Description.ContentMarkdown,
		}
		logrus.Infof("[controller.CreateProduct] create doc %s for product %s", &doc, product)
		return tx.Create(&doc).Error
	})
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return prodID, nil
}

func UpdateProduct(product *model.ProductModel, input *schema.InputProductSchema) error {
	return model.Transact(func(tx *gorm.DB) error {
		// 1. 更新产品本身
		logrus.Infof("[controller.UpdateProduct] update product: %s", product)
		err := tx.Save(&product).Error
		if err != nil {
			return err
		}

		// 2. 更新文档
		doc, err := product.Description()
		if err != nil {
			return err
		}

		err = portal.Dump(&doc, input.Description)
		if err != nil {
			return err
		}

		if doc.ID == 0 {
			// 新建 doc
			logrus.Infof("[controller.UpdateProduct] create doc %s for product: %s", doc, product)
			docID, err := seqgen.NextID()
			if err != nil {
				return err
			}

			doc.ID = docID
			return tx.Create(&doc).Error
		} else {
			// 更新 doc
			logrus.Infof("[controller.UpdateProduct] update doc %s for product: %s", doc, product)
			return tx.Save(&doc).Error
		}
	})
}

func DeleteProduct(product *model.ProductModel) error {
	return model.Transact(func(tx *gorm.DB) error {
		// 删除文档
		err := tx.Delete(&model.DocumentModel{ID: product.DescDocID}).Error
		if err != nil {
			return err
		}

		// 删除产品自身
		err = tx.Delete(&product).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func PublicProduct(prod *model.ProductModel, isPublic bool) error {
	if bool(prod.IsPublic) == isPublic {
		return nil
	}

	doUpdate := func(isPublic bool, publicAt *time.Time) error {
		err := model.DB.Model(&prod).Updates(
			map[string]interface{}{"is_public": isPublic, "public_at": publicAt}).Error
		if err != nil {
			return err
		}

		return nil
	}

	if isPublic == true {
		now := time.Now()
		return doUpdate(true, &now)
	} else {
		return doUpdate(false, nil)
	}
}
