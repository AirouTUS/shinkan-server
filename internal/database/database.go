package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AirouTUS/shinkan-server/internal/model"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
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
		return nil, errors.WithStack(err)
	}

	return m.category(), nil
}

func (db *ShinkanDatabase) GetCircle(input GetCircleInput) (*model.Circle, error) {
	if err := input.validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	var m Circle
	err := db.Map.SelectOne(&m, fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableCircles), input.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return m.circle(), nil
}

func (db *ShinkanDatabase) ListCirclesCircleTypes(input ListCirclesCircleTypesInput) ([]*model.CirclesCircleTypes, error) {
	if err := input.validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	var m CirclesCircleTypesList
	_, err := db.Map.Select(&m, fmt.Sprintf(
		`SELECT 
				%s.circle_type_id,
				%s.name 
			FROM 
				%s 
			JOIN 
				%s
			ON
				%s.circle_type_id = %s.id
			WHERE
				%s.circle_id = ?
			ORDER BY
				%s.id
			ASC`,
		tableCirclesCircleTypes, tableCircleTypes, tableCirclesCircleTypes, tableCircleTypes, tableCirclesCircleTypes, tableCircleTypes, tableCirclesCircleTypes, tableCircleTypes),
		input.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return m.circlesCircleTypes(), nil
}
