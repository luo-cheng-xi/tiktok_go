package terrs

import (
	"errors"
)

// TError 错误信息类，事实上更加接近为前端提供信息的类
type TError struct {
	Code  StatusCode
	Cause error
}

func (t TError) Error() string {
	return t.Cause.Error()
}

// Eq 判断错误是否是Error类型,以及错误是否与该Error相等
func (t TError) Eq(target error) bool {
	switch target.(type) {
	case TError:
		if t.Error() == target.Error() {
			return true
		}
		return false
	}
	return false
}

// 错误类型
var (
	// ErrParamInvalid 无效的参数
	ErrParamInvalid = TError{UNAUTHENTICATED, errors.New("param invalid")}

	// ErrFileInvalid 无效的文件数据
	ErrFileInvalid = TError{INVALID_ARGUMENT, errors.New("file invalid")}

	// ErrTokenInvalid 无效的用户鉴权
	ErrTokenInvalid = TError{UNAUTHENTICATED, errors.New("token invalid")}
	// ErrUserNotFound 无法找到符合要求的用户
	ErrUserNotFound = TError{OK, errors.New("user not found")}
	// ErrUsernameRegistered 用户名已经被注册
	ErrUsernameRegistered = TError{ALREADY_EXISTS, errors.New("username registered")}
	// ErrUsernameTooLong 用户名过长
	ErrUsernameTooLong = TError{INVALID_ARGUMENT, errors.New("username too long")}

	// ErrPasswordWrong 密码错误
	ErrPasswordWrong = TError{INVALID_ARGUMENT, errors.New("password wrong")}
	// ErrPasswordTooLong 密码过长
	ErrPasswordTooLong = TError{INVALID_ARGUMENT, errors.New("password too long")}
	// ErrPasswordTooShort 密码过短
	ErrPasswordTooShort = TError{INVALID_ARGUMENT, errors.New("password too short")}

	// ErrInternal 服务器内部错误
	ErrInternal = TError{INTERNAL, errors.New("internal")}
	// ErrUnknown 未知错误
	ErrUnknown = TError{UNKNOWN, errors.New("unknown")}
)
