package gobatis

import "database/sql"

type DBType string

const (
	DBTypeMySQL      DBType = "mysql"
	DBTypePostgres   DBType = "postgres"
	DBTypeClickhouse DBType = "clickhouse"
)

type GoBatisDB struct {
	db     *sql.DB
	dbType DBType
}

func NewGoBatisDB(dbType DBType, db *sql.DB) *GoBatisDB {
	return &GoBatisDB{
		db:     db,
		dbType: dbType,
	}
}
