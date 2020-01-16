package schema

import "github.com/ifaceless/portal/field"

type OutputDocumentSchema struct {
	ID              string           `json:"id"`
	TOC             []string         `json:"toc"`
	ContentHTML     string           `json:"content_html"`
	ContentMarkdown string           `json:"content_markdown"`
	CreatedAt       *field.Timestamp `json:"created_at,omitempty"`
	UpdatedAt       *field.Timestamp `json:"updated_at,omitempty"`
}

type InputDocumentSchema struct {
	TOC             []string `json:"toc" validate:"required,max=10"`
	ContentHTML     string   `json:"content_html" validate:"max=20480"`
	ContentMarkdown string   `json:"content_markdown" validate:"max=20480"`
}
