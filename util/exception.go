package util

import (
	"fmt"
	"github.com/pkg/errors"
)

type ExceptionInterface interface {
	GetErrCode() int
	GetErrMsg() string
}

type BaseException struct {
	ErrCode int
	ErrMsg  string
}

// Error 异常错误
func (e *BaseException) Error() string {
	return fmt.Sprintf("errCode:%d, errMsg:%s", e.GetErrCode(), e.GetErrMsg())
}

// Code 异常码
func (e *BaseException) GetErrCode() int {
	return e.ErrCode
}

// Msg 异常码
func (e *BaseException) GetErrMsg() string {
	return e.ErrMsg
}

// NewError
func NewError(errCode int, errMsg string) error {
	err := &BaseException{}
	err.ErrCode = errCode
	err.ErrMsg = errMsg
	return errors.Wrap(err, errMsg)
}
