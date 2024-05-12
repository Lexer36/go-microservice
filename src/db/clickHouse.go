package databases

import (
	"database/sql"
	"testProj/src/common"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouseDB struct {
	Db *sql.DB
}

func (db *ClickHouseDB) Init() error {
	// инициализируем соединение с бд
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{common.Config.DbAddr},
		Auth: clickhouse.Auth{
			Database: common.Config.DbName,
			Username: common.Config.DbUser,
			Password: common.Config.DbPassword,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: time.Second * 30,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
	})
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)

	db.Db = conn
	Db = *db
	return conn.Ping()
}

// закрытие соединения
func (db *ClickHouseDB) Close() {
	if db.Db != nil {
		db.Db.Close()
	}
}
