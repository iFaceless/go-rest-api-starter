package resource

import (
	"fmt"

	"github.com/pkg/errors"
)

func errResourceNotFound(kind, rname string) error {
	return errors.New(fmt.Sprintf("%s resource '%s' not found", kind, rname))
}
