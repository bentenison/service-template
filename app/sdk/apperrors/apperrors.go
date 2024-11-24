package apperrors

import (
	"errors"
	"fmt"
	"runtime"
)

type ErrCode struct {
	value int
}
type Errorc struct {
	Code     ErrCode `json:"code"`
	Message  string  `json:"message"`
	FuncName string  `json:"-"`
	FileName string  `json:"-"`
}

// Value returns the integer value of the error code.
func (ec ErrCode) Value() int {
	return ec.value
}

// String returns the string representation of the error code.
func (ec ErrCode) String() string {
	return codeNames[ec]
}

func New(code ErrCode, err error) *Errorc {
	pc, file, line, _ := runtime.Caller(1)
	return &Errorc{
		Code:     code,
		Message:  err.Error(),
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fmt.Sprintf("%s:%d", file, line),
	}
}

func (e *Errorc) Errorc() string {
	return e.Message
}

// NewError checks for an Errorc in the error interface value. If it doesn't
// exist, will create one from the error.
func NewError(err error) *Errorc {
	var errsErr *Errorc
	if errors.As(err, &errsErr) {
		return errsErr
	}

	return New(Internal, err)
}
