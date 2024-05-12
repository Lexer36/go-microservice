package databases

import "database/sql"

// глобальное соединение sql
var (
	Db ClickHouseDB
)

func DbSession() *sql.DB {
	// если соединение потеряно - создаем новое
	if Db.Db == nil {
		db := &ClickHouseDB{}
		db.Init()
	}
	return Db.Db
}
