package resp_code

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// 客户端请求错误
const (
	EmailAlreadyExist = 4000 + iota
	UsernameAlreadyExist
	UsernameDoesNotExist
	InvalidateEmailFormat
	ErrorVerCode
	ExpiredVerCode
	NotExistUsername
	NotExistUserID
	ErrorPwd
	ProblemAlreadyExist
	ProblemNotExist
	UnsupportedLanguage
	SecretError
	UserIDAlreadyExist
	NeedObtainVerificationCode
	CategoryIsNotEmpty
	UserAlreadyRoot
)
