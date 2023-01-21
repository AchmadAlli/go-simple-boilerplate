package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/api"
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/utils"
	"github.com/achmadAlli/go-simple-boilerplate/config"
	"github.com/achmadAlli/go-simple-boilerplate/infrastructures/database"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	loadDatabase()

	e := echo.New()

	// register error handler
	utils.ListenAppError(e)

	// register the route
	api.RegisterRoute(e)

	go func() {
		host := fmt.Sprintf("%s:%d", config.GetEnv().Address, config.GetEnv().Port)
		if err := e.Start(host); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func loadDatabase() {
	config.LoadDBConfig()
	userDB := database.NewMysqlInstance(config.GetEnv().MySQLEnv)

	config.DBConfig.SetUserDB(userDB)
}
