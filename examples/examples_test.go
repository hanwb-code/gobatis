package examples

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql" // 引入驱动
	"github.com/hanwbcode/gobatis"     // 引入gobatis
)

func TestDemo(t *testing.T) {
	// 初始化db
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8")
	dbs := make(map[string]*gobatis.GoBatisDB)
	dbs["ds1"] = gobatis.NewGoBatisDB(gobatis.DBTypeMySQL, db)

	option := gobatis.NewDBOption().
		DB(dbs).
		ShowSQL(true).
		Mappers([]string{"mapper/userMapper.xml"})

	gobatis.Init(option)

	// 获取数据源，参数为数据源名称，如：ds1
	gb := gobatis.Get("ds1")

	//传入id查询Map
	mapRes := make(map[string]interface{})
	_, err := gb.Select("userMapper.findMapById", map[string]interface{}{"id": 1})(mapRes)
	fmt.Println("userMapper.findMapById-->", mapRes, err)

	//mapRes2 := make(map[string]interface{})
	//_, err = gb.SelectContext(context.TODO(), "userMapper.findMapById", map[string]interface{}{"id": 4})(mapRes2)
	//fmt.Println("userMapper.findMapById-->", mapRes2, err)
}
