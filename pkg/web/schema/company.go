package schema

import (
	"github.com/ifaceless/go-starter/pkg/util/pic"
)

type OutputCompanySchema struct {
	ID       string        `json:"id"`
	URL      string        `json:"url"`
	Artworks *pic.Pictures `json:"artwork_urls"`
	Title    string        `json:"title"`
	Intro    string        `json:"intro"`
}
