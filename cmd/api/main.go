package main

import (
	"fmt"
	"user-management/internal/config"
	"user-management/internal/database"
	"user-management/internal/health"
	"user-management/internal/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BUG: why Logger.Info does not work?

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cfg := config.Load()
	gdb, err := database.NewPostgres(cfg)
	fmt.Println("AFTER DB INIT")
	if err != nil {
		e.Logger.Fatal(err)
	}
	err = gdb.AutoMigrate(&user.User{})
	if err != nil {
		e.Logger.Fatal("Migration failed", err)
	}
	e.Logger.Info("Database connected successfully")

	userService := user.NweService(user.Service{DB: gdb})
	userHandler := user.Handler(user.Handler{Service: userService})
	userHandler.SetRoute(e.Group("/users"))

	healthHander := health.NewHandler(health.Handler{})
	healthHander.SetRoute(e.Group("/health"))

	e.Logger.Info("Database migrated successfully")
	e.Logger.Fatal(e.Start(":" + cfg.AppPort))
}
