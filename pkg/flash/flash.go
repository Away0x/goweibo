// Package flash : 模仿 beego flash 的实现
package flash

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// FlashSeparator -
	FlashSeparator = "|"
	// FlashInContextAndCookieKeyName : gin context keys 中 flash 的 key name (cookie 中也是)
	FlashInContextAndCookieKeyName = "flash"
)

// FlashData -
type FlashData struct {
	KeyName string // gin context keys 中的 name (cookie 中也是)
	Data    map[string]string
}

// NewFlashByName - 可取其他 name 的 flash store，可用于其他需要闪存的地方
func NewFlashByName(keyName string) *FlashData {
	return &FlashData{
		KeyName: keyName,
		Data:    make(map[string]string),
	}
}

// NewFlash - flash store
func NewFlash() *FlashData {
	return &FlashData{
		KeyName: FlashInContextAndCookieKeyName,
		Data:    make(map[string]string),
	}
}

// Set message to flash
func (fd *FlashData) Set(key string, msg string, args ...interface{}) {
	if len(args) == 0 {
		fd.Data[key] = msg
	} else {
		fd.Data[key] = fmt.Sprintf(msg, args...)
	}
}

// Success flash
func (fd *FlashData) Success(msg string, args ...interface{}) {
	if len(args) == 0 {
		fd.Data["success"] = msg
	} else {
		fd.Data["success"] = fmt.Sprintf(msg, args...)
	}
}

// Info flash
func (fd *FlashData) Info(msg string, args ...interface{}) {
	if len(args) == 0 {
		fd.Data["info"] = msg
	} else {
		fd.Data["info"] = fmt.Sprintf(msg, args...)
	}
}

// Warning flash
func (fd *FlashData) Warning(msg string, args ...interface{}) {
	if len(args) == 0 {
		fd.Data["warning"] = msg
	} else {
		fd.Data["warning"] = fmt.Sprintf(msg, args...)
	}
}

// Danger falsh
func (fd *FlashData) Danger(msg string, args ...interface{}) {
	if len(args) == 0 {
		fd.Data["danger"] = msg
	} else {
		fd.Data["danger"] = fmt.Sprintf(msg, args...)
	}
}

// Save flash
func (fd *FlashData) Save(c *gin.Context) {
	fd.save(c, FlashInContextAndCookieKeyName)
}

// Read 从 request 中的 cookie 里解析出 flash 数据
func Read(c *gin.Context) *FlashData {
	return read(c, FlashInContextAndCookieKeyName)
}

// NewSuccessFlash : 新建一条 success flash，并保存
func NewSuccessFlash(c *gin.Context, msg string, args ...interface{}) {
	f := NewFlash()
	f.Success(msg, args...)
	f.Save(c)
}

// NewInfoFlash : 新建一条 info flash，并保存
func NewInfoFlash(c *gin.Context, msg string, args ...interface{}) {
	f := NewFlash()
	f.Info(msg, args...)
	f.Save(c)
}

// NewWarningFlash : 新建一条 warning flash，并保存
func NewWarningFlash(c *gin.Context, msg string, args ...interface{}) {
	f := NewFlash()
	f.Warning(msg, args...)
	f.Save(c)
}

// NewDangerFlash : 新建一条 danger flash，并保存
func NewDangerFlash(c *gin.Context, msg string, args ...interface{}) {
	f := NewFlash()
	f.Danger(msg, args...)
	f.Save(c)
}

// ------------------ private
// 将 flash 数据保存到 gin context keys 中和 cookie 中
func (fd *FlashData) save(c *gin.Context, keyName string) {
	c.Keys[keyName] = fd.Data

	var flashValue string
	for key, value := range fd.Data {
		flashValue += "\x00" + key + "\x23" + FlashSeparator + "\x23" + value + "\x00"
	}
	c.SetCookie(keyName, flashValue, 0, "/", "", false, true)
}

// 从 request 中的 cookie 里解析出 flash 数据
func read(c *gin.Context, keyName string) *FlashData {
	flash := NewFlashByName(keyName)
	if cookie, err := c.Request.Cookie(keyName); err == nil {
		v, _ := url.QueryUnescape(cookie.Value)
		vals := strings.Split(v, "\x00")
		for _, v := range vals {
			if len(v) > 0 {
				kv := strings.Split(v, "\x23"+FlashSeparator+"\x23")
				if len(kv) == 2 {
					flash.Data[kv[0]] = kv[1]
				}
			}
		}
		// 读取一次即删除 (beego 里 flash 的实现方式)
		// github.com/tommy351/gin-sessions 的实现方式是，每次 save 都会保存替换所有 session，
		//    所以读取 flash 时，将 flash 从 session 对象中 delete 掉再 save 即可
		c.SetCookie(keyName, "", -1, "/", "", false, true)
	}
	c.Keys[keyName] = flash.Data
	return flash
}
