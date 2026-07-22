package sdk

import (
	"errors"
	"fmt"
)

var ErrObjectNotFound = errors.New("object not found")

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *Error) Is(target error) bool {
	return e != nil && target == ErrObjectNotFound && e.Code == 430004
}
