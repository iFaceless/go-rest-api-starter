package model

import (
	"fmt"
	"time"

	"github.com/ifaceless/go-starter/pkg/util/pic"
)

type CompanyModel struct {
	ID        int64
	Title     string
	Intro     string
	Artworks  *pic.Pictures
	URL       string `gorm:"column:url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *CompanyModel) TableName() string {
	return "company"
}

func (c *CompanyModel) String() string {
	return fmt.Sprintf("<CompanyModel id=%d, title='%s'>", c.ID, c.Title)
}
