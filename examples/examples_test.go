package examples

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql" // 引入驱动
	"github.com/hanwbcode/gobatis"     // 引入gobatis
)

func TestDemo(t *testing.T) {
	var res interface{}

	res = map[string]interface{}{}
	fmt.Println(gobatis.GetResultType(res)) // 输出: map

	res = []map[string]interface{}{}
	fmt.Println(gobatis.GetResultType(res)) // 输出: maps

	res = struct{}{}
	fmt.Println(gobatis.GetResultType(res)) // 输出: struct

	res = []struct{}{}
	fmt.Println(gobatis.GetResultType(res)) // 输出: structs

	res = []interface{}{}
	fmt.Println(gobatis.GetResultType(res)) // 输出: slice

	res = 42
	fmt.Println(gobatis.GetResultType(res)) // 输出: value
}

//func TestDemo(t *testing.T) {
//	// 初始化db
//	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cicd?charset=utf8")
//	dbs := make(map[string]*gobatis.GoBatisDB)
//	dbs["ds1"] = gobatis.NewGoBatisDB(gobatis.DBTypeMySQL, db)
//
//	option := gobatis.NewDBOption().
//		DB(dbs).
//		ShowSQL(true).
//		Mappers([]string{"mapper/userMapper.xml"})
//
//	gobatis.Init(option)
//
//	// 获取数据源，参数为数据源名称，如：ds1
//	gb := gobatis.Get("ds1")
//
//	//传入id查询Map
//	mapRes := make(map[string]interface{})
//
//	_, err := gb.Select("userMapper.findMapById", map[string]interface{}{"id": 1})(mapRes)
//
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("userMapper.findMapById-->", mapRes)
//
//	//mapRes2 := make(map[string]interface{})
//	//_, err = gb.SelectContext(context.TODO(), "userMapper.findMapById", map[string]interface{}{"id": 4})(mapRes2)
//	//fmt.Println("userMapper.findMapById-->", mapRes2, err)
//}
