package main

import (
	"context"
	"github.com/apex/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testDealls/config"
	database "testDealls/dabatase"
	"testDealls/handler"
	"testDealls/repository"
	"testDealls/server"
	"testDealls/service"
	"testDealls/utils"
)

var (
	unhealthy bool
	mut       sync.RWMutex
)

func main() {
	// create context & handle termination signal
	ctx, cancel := context.WithCancel(context.Background())

	exitc := make(chan os.Signal, 1)
	signal.Notify(exitc, os.Interrupt, syscall.SIGTERM)

	// load config
	cfg, err := config.LoadServiceConfig(ctx)
	utils.JWTSecret = cfg.Secret
	if err != nil {
		log.WithError(err).Fatal("can't load config")
	}

	// connect to Mysql
	dbConn := database.Connect(&cfg.MySql)
	defer database.Disconnect(dbConn)

	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)

	// start HTTP server
	e := server.Start(&cfg.HttpServer)
	defer server.Stop(e)

	// handler
	handler.NewUserHandler(e, userService)
	handler.NewHealthCheckHandler(e, dbConn, &unhealthy)

	sig := <-exitc
	log.Infof("received %s signal, exiting now", sig)

	defer cancel()
	setUnhealthy()

}

func setUnhealthy() {
	mut.Lock()
	unhealthy = true
	mut.Unlock()
}
