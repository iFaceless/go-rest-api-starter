package util

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func ParseOrderBy(s string, allowedColumns []string) (string, error) {
	in := strings.Split(s, ",")

	isAllowed := func(s string) bool {
		for _, c := range allowedColumns {
			if c == s {
				return true
			}
		}

		return false
	}

	sqlOrderBys := make([]string, 0, len(in))
	for _, rawOrderBy := range in {
		rawOrderBy = strings.TrimSpace(rawOrderBy)
		if rawOrderBy == "" {
			continue
		}

		column := rawOrderBy
		isDesc := false
		switch rawOrderBy[0] {
		case '+':
			column = rawOrderBy[1:]
		case '-':
			isDesc = true
			column = rawOrderBy[1:]
		default:
		}

		if !isAllowed(column) {
			return "", errors.Errorf("order by '%s' is not allowed", column)
		}

		if isDesc {
			sqlOrderBys = append(sqlOrderBys, quote(column)+" DESC")
		} else {
			sqlOrderBys = append(sqlOrderBys, quote(column))
		}
	}

	return strings.Join(sqlOrderBys, ", "), nil
}

func quote(c string) string {
	if strings.HasPrefix(c, "`") {
		return c
	}

	return fmt.Sprintf("`%s`", c)
}
