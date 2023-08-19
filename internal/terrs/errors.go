package terrs

import (
	"errors"
)

type Error struct {
	Code  StatusCode
	cause error
}

func (e Error) Error() string {
	return e.cause.Error()
}

// Eq 判断错误是否是Error类型,以及错误是否与该Error相等
func (e Error) Eq(target error) bool {
	switch target.(type) {
	case Error:
		if e.Error() == target.Error() {
			return true
		}
		return false
	}
	return false
}

// 错误类型
var (
	// ErrParamInvalid 无效的参数
	ErrParamInvalid = Error{UNAUTHENTICATED, errors.New("param invalid")}

	// ErrTokenInvalid 无效的用户鉴权
	ErrTokenInvalid = Error{UNAUTHENTICATED, errors.New("token invalid")}
	// ErrUserNotFound 无法找到符合要求的用户
	ErrUserNotFound = Error{NOT_FOUND, errors.New("user not found")}
	// ErrUsernameRegistered 用户名已经被注册
	ErrUsernameRegistered = Error{ALREADY_EXISTS, errors.New("username registered")}
	// ErrUsernameTooLong 用户名过长
	ErrUsernameTooLong = Error{INVALID_ARGUMENT, errors.New("username too long")}

	// ErrPasswordWrong 密码错误
	ErrPasswordWrong = Error{INVALID_ARGUMENT, errors.New("password wrong")}
	// ErrPasswordTooLong 密码过长
	ErrPasswordTooLong = Error{INVALID_ARGUMENT, errors.New("password too long")}
	// ErrPasswordTooShort 密码过短
	ErrPasswordTooShort = Error{INVALID_ARGUMENT, errors.New("password too short")}

	// ErrInternal 服务器内部错误
	ErrInternal = Error{INTERNAL, errors.New("internal")}
	// ErrUnknown 未知错误
	ErrUnknown = Error{UNKNOWN, errors.New("unknown")}
)
