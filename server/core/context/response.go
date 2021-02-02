package context

import (
  "goweibo/core/constants"
	"goweibo/core/errno"
	"net/http"
)

// RespData 响应数据类型
type RespData map[string]interface{}

// CommonResponse 通用响应
type CommonResponse struct {
	Code    constants.LogicCode `json:"code"`
	Message string              `json:"msg"`
	Data    interface{}         `json:"data,omitempty"`
}

// NewCommonResponse new CommonResponse
func NewCommonResponse(code constants.LogicCode, message string, data interface{}) *CommonResponse {
	return &CommonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewSuccessResponse new success response
func NewSuccessResponse(message string, data interface{}) *CommonResponse {
	return NewCommonResponse(constants.SuccessCode, message, data)
}

// NewErrResponse new error response
func NewErrResponse(e *errno.Errno) *CommonResponse {
	return NewCommonResponse(e.Code, e.Message, nil)
}

// AWSuccessJSON success response
func (c *AppContext) AWSuccessJSON(data interface{}) error {
	return c.JSON(http.StatusOK, NewSuccessResponse("ok", data))
}

// AWErrorJSON error response
func (c *AppContext) AWErrorJSON(e *errno.Errno) error {
	return c.JSON(e.HTTPCode, NewErrResponse(e))
}
