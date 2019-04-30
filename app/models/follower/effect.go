package follower

import (
	"fmt"
	"gin_weibo/database"
	"strconv"
)

// DoFollow 关注
func DoFollow(userID uint, followIDs ...uint) error {
	l := len(followIDs) - 1
	sqlStr := fmt.Sprintf("insert into %s (follower_id, user_id) values ", tableName)
	for i, v := range followIDs {
		sqlStr = sqlStr + fmt.Sprintf("(%d, %d)", userID, v)
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	d := database.DB.Exec(sqlStr)
	return d.Error
}

// DoUnFollow 取消关注
func DoUnFollow(userID uint, followIDs ...uint) error {
	sqlStr := fmt.Sprintf("delete from %s where follower_id = %d and user_id in (", tableName, userID)
	l := len(followIDs) - 1
	for i, v := range followIDs {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + ")"
	d := database.DB.Exec(sqlStr)
	return d.Error
}
