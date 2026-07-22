package sdk

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorIsObjectNotFound(t *testing.T) {
	if !errors.Is(&Error{Code: 430004}, ErrObjectNotFound) {
		t.Fatal("expected 430004 to match ErrObjectNotFound")
	}
	wrapped := fmt.Errorf("get folder info: %w", &Error{Code: 430004})
	if !errors.Is(wrapped, ErrObjectNotFound) {
		t.Fatal("expected wrapped 430004 to match ErrObjectNotFound")
	}
	var sdkErr *Error
	if !errors.As(wrapped, &sdkErr) || sdkErr.Code != 430004 {
		t.Fatal("expected wrapped error to preserve *Error")
	}
	if errors.Is(&Error{Code: 430005}, ErrObjectNotFound) {
		t.Fatal("unexpected object-not-found match")
	}
	if errors.Is(&Error{Code: 430004}, errors.New("object not found")) {
		t.Fatal("unexpected match for a different target")
	}
	var nilErr *Error
	if errors.Is(nilErr, ErrObjectNotFound) {
		t.Fatal("unexpected match for a nil error")
	}
}
