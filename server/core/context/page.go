package context

// CommonPageData 通用的分页数据
type CommonPageData struct {
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageLine int           `json:"line"`
	List     []interface{} `json:"list"`
}

// NewCommonPageData new CommonPageData
func NewCommonPageData(total, page, pageLine int, list []interface{}) *CommonPageData {
	return &CommonPageData{
		Total:    total,
		Page:     page,
		PageLine: pageLine,
		List:     list,
	}
}

// AWPageJSON page response
func (c *AppContext) AWPageJSON(page *CommonPageData) error {
	return c.AWSuccessJSON(page)
}
