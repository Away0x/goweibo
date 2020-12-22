package constants

type (
	// LogicCode logic code
	LogicCode int
)

const (
	// SuccessCode success code
	SuccessCode LogicCode = 0
	// UnknownErrorCode unknown error code
	UnknownErrorCode LogicCode = -1
	// RequestErrorCode request error code
	RequestErrorCode LogicCode = 100
	// ResourceErrorCode resource error code
	ResourceErrorCode LogicCode = 101
	// DatabaseErrorCode database error code
	DatabaseErrorCode LogicCode = 102
	// TokenErrorCode token error code
	TokenErrorCode LogicCode = 103
	// NotFoundErrorCode not found error code
	NotFoundErrorCode LogicCode = 104
)
