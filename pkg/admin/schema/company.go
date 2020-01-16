package schema

import (
	"github.com/ifaceless/go-starter/pkg/util/pic"
	"github.com/ifaceless/portal/field"
)

type OutputCompanySchema struct {
	ID        string           `json:"id"`
	URL       string           `json:"url"`
	Artworks  *pic.Pictures    `json:"artwork_urls" portal:"AUTO_INIT"`
	Title     string           `json:"title"`
	Intro     string           `json:"intro"`
	CreatedAt *field.Timestamp `json:"created_at"`
	UpdatedAt *field.Timestamp `json:"updated_at"`
}

type InputCompanySchema struct {
	URL      string        `json:"url"`
	Artworks *pic.Pictures `json:"artwork_urls"`
	Title    string        `json:"title"`
	Intro    string        `json:"intro"`
}
