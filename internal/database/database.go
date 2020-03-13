package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AirouTUS/shinkan-server/internal/model"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const (
	driverName = "mysql"
	engineName = "InnoDB"

	tableCircles          = "circles"
	tableCircleCategories = "circle_categories"
	tableCircleTypes      = "circle_types"
	tableCircleImages     = "circle_images"

	// 中間テーブル
	tableCirclesCircleTypes = "circles_circle_types"
)

type ShinkanDatabase struct {
	Map *gorp.DbMap
}

func NewDatabase(user, password, host, port, database string) *ShinkanDatabase {
	db, err := sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: engineName, Encoding: "UTF8"}}
	return &ShinkanDatabase{Map: dbMap}
}

func (db *ShinkanDatabase) ListCategory(input ListCategoryInput) ([]*model.Category, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	var m CategoryList
	_, err := db.Map.Select(&m, fmt.Sprintf("SELECT id, name FROM %s ORDER BY id ASC", tableCircleCategories))
	if err != nil {
		return nil, err
	}

	return m.category(), nil
}
