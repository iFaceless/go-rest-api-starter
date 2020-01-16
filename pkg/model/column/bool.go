package column

import (
	"database/sql"
	"database/sql/driver"
)

type Bool bool

func (b *Bool) Value() (driver.Value, error) {
	if bool(*b) == true {
		return 1, nil
	} else {
		return 0, nil
	}
}

func (b *Bool) Scan(v interface{}) error {
	dst := sql.NullBool{}
	err := dst.Scan(v)
	if err != nil {
		return err
	}

	*b = Bool(dst.Bool)
	return nil
}
