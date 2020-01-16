// Package pic
package pic

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/spf13/cast"
)

type Picture struct {
	token string
}

func NewPicture(tokenOrURL string) *Picture {
	return &Picture{token: parseURLToken(tokenOrURL)}
}

func (p *Picture) Value() (driver.Value, error) {
	return p.token, nil
}

func (p *Picture) Scan(v interface{}) error {
	p.token = NewPicture(cast.ToString(v)).token
	return nil
}

func (p *Picture) SetValue(v interface{}) error {
	return p.Scan(v)
}

func (p *Picture) MarshalJSON() ([]byte, error) {
	return json.Marshal(GetFullURL(p.token))
}

func (p *Picture) UnmarshalJSON(v []byte) error {
	var token string
	err := json.Unmarshal(v, &token)
	if err != nil {
		return err
	}

	p.token = parseURLToken(token)
	return nil
}

type Pictures []*Picture

func NewPictures(tokenOrURLs []string) *Pictures {
	pics := make(Pictures, 0, len(tokenOrURLs))
	for _, u := range tokenOrURLs {
		pics = append(pics, NewPicture(u))
	}

	return &pics
}

func (p *Pictures) Value() (driver.Value, error) {
	tokens := make([]string, 0)
	for _, pic := range *p {
		tokens = append(tokens, pic.token)
	}

	return json.Marshal(tokens)
}

func (p *Pictures) Scan(v interface{}) error {
	data, ok := v.([]byte)
	if !ok || len(data) == 0 {
		return nil
	}

	var tokens []string
	err := json.Unmarshal(data, &tokens)
	if err != nil {
		return err
	}

	*p = *NewPictures(tokens)
	return nil
}

func (p *Pictures) SetValue(v interface{}) error {
	return p.Scan(v)
}

func (p *Pictures) MarshalJSON() ([]byte, error) {
	tokens := make([]string, 0)
	for _, item := range *p {
		if item != nil {
			tokens = append(tokens, GetFullURL(item.token))
		}
	}
	return json.Marshal(tokens)
}

func (p *Pictures) UnmarshalJSON(v []byte) error {
	var tokens []string
	err := json.Unmarshal(v, &tokens)
	if err != nil {
		return err
	}

	*p = *NewPictures(tokens)
	return nil
}
