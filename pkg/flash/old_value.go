package flash

import (
	"github.com/gin-gonic/gin"
)

const (
	// OldValueInContextAndCookieKeyName : gin context keys 中的 name (cookie 中也是)
	OldValueInContextAndCookieKeyName = "oldValue"
)

// SaveOldFormValue : 存储上次表单 post 的数据
func SaveOldFormValue(c *gin.Context, obj map[string]string) {
	f := NewFlashByName(OldValueInContextAndCookieKeyName)
	f.Data = obj
	f.save(c, OldValueInContextAndCookieKeyName)
}

// ReadOldFormValue : 读取上次表单 post 的数据
func ReadOldFormValue(c *gin.Context) *FlashData {
	return read(c, OldValueInContextAndCookieKeyName)
}
