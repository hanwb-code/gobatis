package examples

import "github.com/hanwbcode/gobatis"

// 实体结构示例， tag：field为数据库对应字段名称
type User struct {
	Id    gobatis.NullInt64  `field:"id"`
	Name  gobatis.NullString `field:"name"`
	Email gobatis.NullString `field:"email"`
	CrtTm gobatis.NullTime   `field:"crtTm"`
}
