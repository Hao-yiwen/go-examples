package errors

import (
	"errors"
	"fmt"
)

// 通用错误定义
var (
	ErrNotFound       = errors.New("resource not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrInternal       = errors.New("internal error")
	ErrAlreadyExists  = errors.New("resource already exists")
	ErrValidation     = errors.New("validation error")
)

// AppError 应用错误
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error 实现error接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap 支持errors.Is和errors.As
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError 创建应用错误
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Wrap 包装错误
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// Wrapf 格式化包装错误
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}

// Is 检查错误类型
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As 错误类型断言
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// New 创建新错误
func New(message string) error {
	return errors.New(message)
}

// Errorf 格式化创建错误
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// HTTP错误码
const (
	CodeSuccess         = 0
	CodeBadRequest      = 400
	CodeUnauthorized    = 401
	CodeForbidden       = 403
	CodeNotFound        = 404
	CodeConflict        = 409
	CodeInternalError   = 500
)

// 预定义应用错误
func ErrBadRequest(message string) *AppError {
	return NewAppError(CodeBadRequest, message, nil)
}

func ErrUnauthorizedError(message string) *AppError {
	return NewAppError(CodeUnauthorized, message, nil)
}

func ErrForbiddenError(message string) *AppError {
	return NewAppError(CodeForbidden, message, nil)
}

func ErrNotFoundError(message string) *AppError {
	return NewAppError(CodeNotFound, message, nil)
}

func ErrConflict(message string) *AppError {
	return NewAppError(CodeConflict, message, nil)
}

func ErrInternalError(message string) *AppError {
	return NewAppError(CodeInternalError, message, nil)
}
