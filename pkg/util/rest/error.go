package rest

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	StatusCode int         `json:"-"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Name       string      `json:"name"`
	Data       interface{} `json:"data,omitempty"`
}

type Option func(e *APIError)

func Message(s string) Option {
	return func(e *APIError) {
		if s == "" {
			s = "Bad Request"
		}
		e.Message = s
	}
}

func Code(v int) Option {
	return func(e *APIError) {
		e.Code = v
	}
}

func StatusCode(v int) Option {
	return func(e *APIError) {
		e.StatusCode = v
	}
}

func Name(n string) Option {
	return func(e *APIError) {
		if n == "" {
			n = "BadRequest"
		}
		e.Name = n
	}
}

func Data(d interface{}) Option {
	return func(e *APIError) {
		e.Data = d
	}
}

func New(opts ...Option) *APIError {
	err := &APIError{
		StatusCode: http.StatusBadRequest,
		Code:       400,
		Name:       "BadRequest",
		Message:    "Bad Request",
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError(status=%d, code=%d, name='%s', message='%s')",
		e.StatusCode, e.Code, e.Name, e.Message)
}

func abortWithError(c *gin.Context, err error) {
	apiErr, ok := err.(*APIError)
	if ok {
		c.AbortWithStatusJSON(apiErr.StatusCode, gin.H{"error": apiErr})
	} else {
		bad := BadRequest(err)
		c.AbortWithStatusJSON(bad.StatusCode, gin.H{"error": bad})
	}
}

//
// Some frequently used errors
//
func InternalServerError(err error, opts ...Option) *APIError {
	if err == nil {
		err = errors.New("Internal server error")
	}

	defaultOpts := []Option{StatusCode(http.StatusInternalServerError), Message(err.Error()), Code(http.StatusInternalServerError), Name("InternalServerError")}
	return New(append(defaultOpts, opts...)...)
}

func ResourceNotFound(err error, opts ...Option) *APIError {
	if err == nil {
		err = errors.New("Resource not found error")
	}

	defaultOpts := []Option{StatusCode(http.StatusNotFound), Message(err.Error()), Code(http.StatusNotFound), Name("ResourceNotFound")}
	return New(append(defaultOpts, opts...)...)
}

func Unauthorized(err error, opts ...Option) *APIError {
	if err == nil {
		err = errors.New("Unauthorized error")
	}

	defaultOpts := []Option{StatusCode(http.StatusUnauthorized), Message(err.Error()), Code(http.StatusUnauthorized), Name("Unauthorized")}
	return New(append(defaultOpts, opts...)...)
}

func RequestForbidden(err error, opts ...Option) *APIError {
	if err == nil {
		err = errors.New("Request forbidden error")
	}

	defaultOpts := []Option{StatusCode(http.StatusForbidden), Message(err.Error()), Code(http.StatusForbidden), Name("RequestForbidden")}
	return New(append(defaultOpts, opts...)...)
}

func BadRequest(err error, opts ...Option) *APIError {
	if err == nil {
		err = errors.New("Bad request error")
	}

	defaultOpts := []Option{StatusCode(http.StatusBadRequest), Message(err.Error()), Code(http.StatusBadRequest), Name("BadRequest")}
	return New(append(defaultOpts, opts...)...)
}
