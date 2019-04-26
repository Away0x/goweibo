package pagination

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

type paginationRenderData struct {
	URL          string // 分页的 root url
	CurrentPage  int    // 当前页码
	OnFirstPage  bool   // 是否在第一页
	HasMorePages bool   // 是否有更多页
	Elements     []int  // 页码

	PreviousButtonText string // 前一页按钮文本
	PreviousPageIndex  int    // 前一页按钮的页码
	NextButtonText     string // 后一页按钮文本
	NextPageIndex      int    // 后一页按钮的页码
}

// CreatePaginationFillToTplData : 生成分页模板所需的数据
func CreatePaginationFillToTplData(c *gin.Context, pageQueryKeyName string, currentPage, totalPage int, otherData map[string]interface{}) map[string]interface{} {
	queryValues := url.Values{}
	for k, v := range c.Request.URL.Query() {
		if k != pageQueryKeyName {
			queryValues.Add(k, v[0])
		}
	}
	query := queryValues.Encode()
	if query != "" {
		query = query + "&"
	}

	pageData := paginationRenderData{
		URL:          c.Request.URL.Path + "?" + query + pageQueryKeyName + "=",
		CurrentPage:  currentPage,
		OnFirstPage:  currentPage == 1,
		HasMorePages: currentPage != totalPage,
		Elements:     countStartAndEndPageIndex(currentPage, totalPage, 3),

		PreviousButtonText: "前一页",
		PreviousPageIndex:  currentPage - 1,
		NextButtonText:     "下一页",
		NextPageIndex:      currentPage + 1,
	}

	otherData["pagination"] = pageData
	return otherData
}

// 返回一个区间数组，供生成区间页码按钮
// baseOnCurrentPageButtonOffset: 前后有多少个按钮
func countStartAndEndPageIndex(currentPage, totalPage, baseOnCurrentPageButtonOffset int) []int {
	howMuchPageButtons := baseOnCurrentPageButtonOffset*2 + 1
	startPage := 1
	endPage := 1
	result := make([]int, 0)

	if currentPage > baseOnCurrentPageButtonOffset {
		// 当前页码大于偏移量，则起始按钮为 当前页码 - 偏移量

		startPage = currentPage - baseOnCurrentPageButtonOffset
		if totalPage > (currentPage + baseOnCurrentPageButtonOffset) {
			endPage = currentPage + baseOnCurrentPageButtonOffset
		} else {
			endPage = totalPage
		}
	} else {
		// 当前页码小于偏移量

		startPage = 1
		if totalPage > howMuchPageButtons {
			endPage = howMuchPageButtons
		} else {
			endPage = totalPage
		}
	}

	if (currentPage + baseOnCurrentPageButtonOffset) > totalPage {
		startPage = startPage - (currentPage + baseOnCurrentPageButtonOffset - endPage)
	}

	if startPage <= 0 {
		startPage = 1
	}

	for i := startPage; i <= endPage; i++ {
		result = append(result, i)
	}

	return result
}
