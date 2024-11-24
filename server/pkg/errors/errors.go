package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	Code    int
	Massage string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Massage)
}

func As(err error) (*Error, bool) {
	if err == nil {
		return nil, false
	}
	var merr *Error
	if errors.As(err, &merr) {
		return merr, true
	}
	return nil, false
}

func InternalServer(format string, v ...any) *Error {
	return &Error{
		Code:    500,
		Massage: fmt.Sprintf(format, v...),
	}
}

func InvalidInput(format string, v ...any) *Error {
	return &Error{
		Code:    400,
		Massage: fmt.Sprintf(format, v...),
	}
}

func AuthFailed(format string, v ...any) *Error {
	return &Error{
		Code:    401,
		Massage: fmt.Sprintf(format, v...),
	}
}
