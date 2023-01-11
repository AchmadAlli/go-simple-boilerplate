package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/achmadAlli/go-simple-boilerplate/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	e := echo.New()

	e.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "server is healthy")
	})

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
