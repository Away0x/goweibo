package context

import "goweibo/core/pkg/session"

// AWSession 获取 session 实例
func (c *AppContext) AWSession() session.Session {
	return session.Default(c)
}

// AWSessionSet 设置 session
func (c *AppContext) AWSessionSet(key, val string) {
	session.Set(c, key, val)
}

// AWSessionGet 获取 session
func (c *AppContext) AWSessionGet(key string) string {
	return session.Get(c, key)
}

// AWSessionDelete 删除 session
func (c *AppContext) AWSessionDelete(key string) {
	session.Delete(c, key)
}
