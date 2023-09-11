package gobatis

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"testing"
)

type TUser2 struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

func TestGoBatisWithDBCli(t *testing.T) {

	db, _ := sql.Open("clickhouse", "clickhouse://root:12345@127.0.0.1:9000/cloud_cost_analysis_business")
	dbs := make(map[string]*GoBatisDB)
	dbs["ds"] = NewGoBatisDB(DBTypeClickhouse, db)

	option := NewDBOption().
		DB(dbs).
		ShowSQL(true).
		Mappers([]string{"examples/mapper/userMapper.xml"})
	Init(option)

	if nil == conf {
		LOG.Info("db config == nil")
		return
	}

	gb := Get("ds")

	//var result *TUser2
	//_, err := gb.Select("userMapper.findById", map[string]interface{}{
	//	"id": 1,
	//})(&result)
	//
	//fmt.Println("result:", result, "err:", err)

	//param := &TUser2{}
	//
	//res := make([]*TUser2, 0)
	//
	//_, err := gb.Select(`
	//	<select id="list" resultType="structs">
	//		SELECT
	//			id,
	//			name
	//		FROM
	//			users
	//		<where>
	//			<if test="Name != nil and Name != ''">and name = #{Name}</if>
	//		</where>
	//	</select>
	//`, param)(&res)
	//
	//for _, re := range res {
	//	fmt.Printf("list %+v\n", re)
	//}

	uu := &Users{
		Id:   1,
		Name: "1993",
	}

	// test set
	affected, err := gb.Update("userMapper.updateByCond", uu)

	fmt.Println("updateByCond:", affected, err)

}
