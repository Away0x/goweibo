package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIntParam : 从 path params 中获取 int 参数
// http://a.com/xx/1 => 获取到 int 1
func GetIntParam(c *gin.Context, key string) (int, error) {
	i, err := strconv.Atoi(c.Param(key))
	if err != nil {
		return 0, err
	}

	return i, nil
}
