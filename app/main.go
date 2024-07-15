package main

import (
	"fmt"
	"os"

	"github.com/q4Zar/go-rest-api/database/repository"
	// seeders "github.com/q4Zar/go-rest-api/database/seed"
	"github.com/q4Zar/go-rest-api/http/route"
	"github.com/q4Zar/go-rest-api/service/currency"
	"github.com/q4Zar/go-rest-api/service/user"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
	_ "goyave.dev/goyave/v5/database/dialect/postgres"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	opts := goyave.Options{
		Config: cfg,
	}


	server, err := goyave.New(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	server.Logger.Info("Registering hooks")
	server.RegisterSignalHook()

	server.RegisterStartupHook(func(s *goyave.Server) {
		s.Logger.Info("Server is listening", "host", s.Host())
	})

	server.RegisterShutdownHook(func(s *goyave.Server) {
		s.Logger.Info("Server is shutting down")
	})

	registerServices(server)

	server.Logger.Info("Registering routes")
	server.RegisterRoutes(route.Register)

	if err := server.Start(); err != nil {
		server.Logger.Error(err)
		os.Exit(2)
	}
}

func registerServices(server *goyave.Server) {
	server.Logger.Info("Registering services")

	session := session.GORM(server.DB(), nil)

	userRepo := repository.NewUser(server.DB())
	currencyRepo := repository.NewCurrency(server.DB())

	server.RegisterService(user.NewService(session, server.Logger, userRepo))
	server.RegisterService(currency.NewService(session, currencyRepo))
}