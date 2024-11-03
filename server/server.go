package server

import (
	"context"
	"errors"
	"github.com/apex/log"
	"github.com/labstack/echo/v4"
	"net/http"
	"testDealls/config"
	"time"
)

func Start(c *config.HttpServerConfig) *echo.Echo {

	e := echo.New()

	go func() {
		log.Warnf("[server] Starting the apps on port %s \n", c.Port)
		if err := e.Start(c.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[server] Shutting the apps... [%s]", err)
		}
	}()

	return e

}

func Stop(e *echo.Echo) {

	log.Info("stopping HTTP server")
	log.Info("wait 30 second before calling server shutdown")
	time.Sleep(time.Second * 30)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
