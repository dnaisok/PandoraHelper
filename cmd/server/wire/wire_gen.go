// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"PandoraHelper/internal/handler"
	"PandoraHelper/internal/repository"
	"PandoraHelper/internal/server"
	"PandoraHelper/internal/service"
	"PandoraHelper/pkg/app"
	"PandoraHelper/pkg/jwt"
	"PandoraHelper/pkg/log"
	"PandoraHelper/pkg/server/http"
	"PandoraHelper/pkg/sid"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	db := repository.NewDB(viperViper, logger)
	repositoryRepository := repository.NewRepository(logger, db)
	transaction := repository.NewTransaction(repositoryRepository)
	sidSid := sid.NewSid()
	serviceService := service.NewService(transaction, logger, sidSid, jwtJWT)
	userService := service.NewUserService(serviceService, viperViper)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	shareRepository := repository.NewShareRepository(repositoryRepository)
	shareService := service.NewShareService(serviceService, shareRepository)
	shareHandler := handler.NewShareHandler(handlerHandler, shareService)
	accountRepository := repository.NewAccountRepository(repositoryRepository)
	accountService := service.NewAccountService(serviceService, accountRepository)
	accountHandler := handler.NewAccountHandler(handlerHandler, accountService)
	httpServer := server.NewHTTPServer(logger, viperViper, jwtJWT, userHandler, shareHandler, accountHandler)
	job := server.NewJob(logger)
	appApp := newApp(httpServer, job)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRepository, repository.NewTransaction, repository.NewAccountRepository, repository.NewShareRepository)

var serviceSet = wire.NewSet(service.NewService, service.NewUserService, service.NewAccountService, service.NewShareService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler, handler.NewShareHandler, handler.NewAccountHandler)

var serverSet = wire.NewSet(server.NewHTTPServer, server.NewJob)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(app.WithServer(httpServer, job), app.WithName("demo-server"))
}
