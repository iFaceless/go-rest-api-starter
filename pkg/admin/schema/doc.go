package schema

import "github.com/ifaceless/portal/field"

type OutputDocumentSchema struct {
	ID              *string          `json:"id"`
	ContentHTML     *string          `json:"content_html,omitempty"`
	ContentMarkdown *string          `json:"content_markdown,omitempty"`
	CreatedAt       *field.Timestamp `json:"created_at,omitempty"`
	UpdatedAt       *field.Timestamp `json:"updated_at,omitempty"`
}

type InputDocumentSchema struct {
	ContentHTML     string `json:"content_html" validate:"max=20480"`
	ContentMarkdown string `json:"content_markdown" validate:"max=20480"`
}
