// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"search-product/app/search/internal/biz"
	"search-product/app/search/internal/conf"
	"search-product/app/search/internal/data"
	"search-product/app/search/internal/server"
	"search-product/app/search/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	searcherRepo := data.NewSearcherRepo(dataData, logger)
	searcherUsecase := biz.NewSearcherUsecase(searcherRepo, logger)
	searchService := service.NewSearchService(searcherUsecase)
	grpcServer := server.NewGRPCServer(confServer, searchService, logger)
	httpServer := server.NewHTTPServer(confServer, searchService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
