package examples

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql" // 引入驱动
	"github.com/hanwbcode/gobatis"     // 引入gobatis
)

func Test_Clickhouse(t *testing.T) {
	// 初始化db
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	dbs := make(map[string]*gobatis.GoBatisDB)
	dbs["ds1"] = gobatis.NewGoBatisDB(gobatis.DBTypeMySQL, db)

	option := gobatis.NewDBOption().
		DB(dbs).
		ShowSQL(true).
		Mappers([]string{"mapper/userMapper.xml"})

	gobatis.Init(option)

	// 获取数据源，参数为数据源名称，如：ds1
	gb := gobatis.Get("ds1")

	mapRes2 := make(map[string]interface{})
	_, err := gb.SelectContext(context.TODO(), "userMapper.findMapById", map[string]interface{}{"id": 4})(mapRes2)
	fmt.Println("userMapper.findMapById-->", mapRes2, err)
}
