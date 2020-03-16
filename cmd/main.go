package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/AirouTUS/shinkan-server/pkg/app/handler"

	"github.com/AirouTUS/shinkan-server/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port = flag.String("port", ":8080", "port to listen")
)

func main() {
	flag.Parse()
	checkEnv()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	dbRepo := database.NewDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))
	apiHandler := handler.NewHandler(dbRepo)

	var err error
	handler.Categories, err = dbRepo.ListCategory(database.ListCategoryInput{})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	e.GET("/categories", apiHandler.ListCategory)
	e.GET("/circles/:id", apiHandler.GetCircle)
	e.GET("/circles", apiHandler.ListCircle)

	basic_username := os.Getenv("SHINKAN_USER_NAME")
	basic_password := os.Getenv("SHINAKN_PASSWORD")
	m := e.Group("/admin", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == basic_username && password == basic_password {
			return true, nil
		}
		return false, nil
	}))
	m.POST("/circle", func(c echo.Context) error {
		return c.NoContent(http.StatusCreated)
	})

	e.Logger.Fatal(e.Start(*port))
}

func checkEnv() {
	switch "" {
	case os.Getenv("MYSQL_USER"):
		log.Println("MYSQL_USER is undefined")
		os.Exit(1)
	case os.Getenv("MYSQL_PASSWORD"):
		log.Println("MYSQL_PASSWORD is undefined")
		os.Exit(1)
	case os.Getenv("MYSQL_HOST"):
		log.Println("MYSQL_HOST is undefined")
		os.Exit(1)
	case os.Getenv("MYSQL_PORT"):
		log.Println("MYSQL_PORT is undefined")
		os.Exit(1)
	case os.Getenv("MYSQL_DATABASE"):
		log.Println("MYSQL_DATABASE is undefined")
		os.Exit(1)
	case os.Getenv("SHINKAN_USER_NAME"):
		log.Println("SHINKAN_USER_NAME is undefined")
		os.Exit(1)
	case os.Getenv("SHINKAN_PASSWORD"):
		log.Println("SHINKAN_PASSWORD is undefined")
		os.Exit(1)
	}
}
