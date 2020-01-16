package model

import (
	"fmt"
	"time"
)

type DocumentModel struct {
	ID              int64
	ContentHTML     string              `gorm:"column:content_html"`
	ContentMarkdown string              `gorm:"column:content_md"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (doc *DocumentModel) TableName() string {
	return "doc"
}

func (doc *DocumentModel) String() string {
	return fmt.Sprintf("<DocumentModel id=%d>", doc.ID)
}
