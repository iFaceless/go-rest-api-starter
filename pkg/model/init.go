package model

import (
	"fmt"
	"net/url"

	"github.com/ifaceless/go-starter/pkg/config"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

func init() {
	db, err := gorm.Open("mysql", getDsn(config.MySQLConfig.Master))
	if err != nil {
		logrus.Panicf("failed to open db with config '%s':%s", config.MySQLConfig.Master, err)
	}
	DB = db
}

func getDsn(rawurl string) string {
	u, err := url.Parse("mysql://" + rawurl)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s@tcp(%s)%s?%s", u.User, u.Host, u.Path, u.RawQuery)
}

// Transact auto commit or rollback on error.
func Transact(fn func(tx *gorm.DB) error) (err error) {
	tx := DB.Begin()
	if tx.Error != nil {
		return errors.WithStack(tx.Error)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.RollbackUnlessCommitted()
			err = errors.New(fmt.Sprintf("%s", p))
		}
	}()

	err = fn(tx)
	if err != nil {
		tx.RollbackUnlessCommitted()
		return errors.WithStack(err)
	}

	return errors.WithStack(tx.Commit().Error)
}
