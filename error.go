package sdk

import (
	"errors"
	"fmt"
)

var (
	ErrObjectNotFound = errors.New("object not found")
	ErrOfflineTaskAlreadyExists = errors.New("offline task already exists")
)

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *Error) Is(target error) bool {
	if e == nil {
		return false
	}
	switch target {
	case ErrObjectNotFound:
		return e.Code == 430004
	case ErrOfflineTaskAlreadyExists:
		return e.Code == 10008
	}
	return false
}
