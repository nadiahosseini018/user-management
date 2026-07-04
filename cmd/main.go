package main

import (
	"fmt"
	"net/http"
	"user-management/internal/config"
	"user-management/internal/database"
	"user-management/internal/models"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Application is running")
	})
	cfg := config.Load()
	db, err := database.NewPostgres(cfg)

	fmt.Println("AFTER DB INIT")

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Info("Database connected successfully")

	_ = db

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		e.Logger.Fatal("Migration failed", err)
	}
	e.Logger.Info("Database migrated successfully")
	e.Logger.Fatal(e.Start(":" + cfg.AppPort))
}
