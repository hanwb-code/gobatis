package clickhouse

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go/v2"
)

const Clickhouse = "clickhouse"

func InitDB(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open(Clickhouse, dsn)
	if err != nil {
		err = fmt.Errorf("Clickhouse connnet error: %s", err)
		return
	}
	return
}
