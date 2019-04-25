package controllers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPageQuery 从 query 中获取有关分页的参数
// xx.com?page=1&pageline=10
func GetPageQuery(c *gin.Context, defaultPageLine, totalCount int) (offset, limit, currentPage, pageTotalCount int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	currentPage = page
	pageline, err := strconv.Atoi(c.Query("pageline"))
	if err != nil {
		pageline = defaultPageLine
	}

	page = page - 1
	if page == 0 {
		offset = 0
	} else {
		offset = page * pageline
	}

	limit = pageline

	pageTotalCount = int(math.Ceil(float64(totalCount) / float64(pageline)))
	if pageTotalCount <= 0 {
		pageTotalCount = 1
	}

	return
}
