package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPageQuery 从 query 中获取有关分页的参数
// xx.com?page=1&pageline=10
func GetPageQuery(c *gin.Context, defaultPageLine int) (offset, limit int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pageline, err := strconv.Atoi(c.Query("pageline"))
	if err != nil {
		pageline = defaultPageLine
	}

	page = page - 1
	if page == 0 {
		offset = 0
	} else {
		offset = (page * pageline) - 1
	}

	limit = pageline

	return
}
