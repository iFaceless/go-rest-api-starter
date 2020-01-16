package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type Context struct {
	*gin.Context
}

func (c *Context) Offset() int {
	return cast.ToInt(c.Query("offset"))
}

func (c *Context) Limit() int {
	limit := cast.ToInt(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	return limit
}

func (c *Context) IntParam(key string) int {
	return cast.ToInt(c.Param(key))
}

type validator interface {
	Validate() error
}

func (c *Context) BindJSON(v interface{}) error {
	err := c.Context.BindJSON(v)
	if err != nil {
		return errors.WithStack(err)
	}

	if validate, ok := v.(validator); ok {
		err = validate.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Context) QueryIntArray(key string) []int {
	array := c.QueryArray(key)
	result := make([]int, 0, len(array))
	for _, item := range array {
		v := cast.ToInt(item)
		if v != 0 {
			result = append(result, cast.ToInt(item))
		}
	}
	return result
}

func (c *Context) QueryWithFallback(key string, fallback string) string {
	r := c.Query(key)
	if r == "" {
		return fallback
	}
	return r
}

func (c *Context) QueryBool(key string) bool {
	return cast.ToBool(c.Query(key))
}
