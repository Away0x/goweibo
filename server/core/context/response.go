package context

import (
	"goweibo/core/errno"
	"net/http"
)

// RespData 响应数据类型
type RespData map[string]interface{}

// CommonResponse 通用响应
type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// NewCommonResponse new CommonResponse
func NewCommonResponse(code int, message string, data interface{}) *CommonResponse {
	return &CommonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewSuccessResponse new success response
func NewSuccessResponse(message string, data interface{}) *CommonResponse {
	return NewCommonResponse(0, message, data)
}

// NewErrResponse new error response
func NewErrResponse(e *errno.Errno) *CommonResponse {
	return NewCommonResponse(e.Code, e.Message, nil)
}

// SuccessResp success response
func (c *AppContext) SuccessResp(data RespData) error {
	return c.JSON(http.StatusOK, NewSuccessResponse("", data))
}

// ErrorResp error response
func (c *AppContext) ErrorResp(e *errno.Errno) error {
	return c.JSON(e.HTTPCode, NewErrResponse(e))
}
