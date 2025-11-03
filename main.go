package main

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	service "article/pkg/modules/article"
	"article/pkg/store/article"
)

func main() {
	initConfig()

	port := viper.GetString("db.article.port")
	userName := viper.GetString("db.article.user")
	password := viper.GetString("db.article.password")
	host := viper.GetString("db.article.host")
	dbname := viper.GetString("db.article.name")

	dbStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	fmt.Printf("Attempting to connect to database at %s:%s\n", host, port)

	// Connect to database
	db, err := sqlx.Connect("postgres", dbStr)
	if err != nil {
		fmt.Printf("failed to connect to db: %v\n", err)
		fmt.Printf("Connection string: postgres://%s:***@%s:%s/%s?sslmode=disable\n", userName, host, port, dbname)
		return
	}
	fmt.Println("Successfully connected to database")

	m, err := migrate.New(
		"file://migrates",
		dbStr)

	if err != nil {
		fmt.Println("failed to make migrate", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("failed to m.Up() in migrate")
		return
	}

	e := echo.New()

	userStore := article.NewArticleRepository(db)
	s := service.NewService(userStore, db, e)

	e.GET("/article/:id", s.Get)
	e.POST("/article", s.Create)

	e.Start(":8000")
}

func initConfig() {
	viper.SetConfigName("testing")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file not found")
	}
}
