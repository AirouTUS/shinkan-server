package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const (
	driverName = "mysql"
	engineName = "InnoDB"
)

type DojoDatabase struct {
	Map *gorp.DbMap
}

func NewDatabase(user, password, host, port, database string) *DojoDatabase {
	db, err := sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: engineName, Encoding: "UTF8"}}
	return &DojoDatabase{Map: dbMap}
}
