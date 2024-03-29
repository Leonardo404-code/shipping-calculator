// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handler

import (
	"database/sql"
	"shipping-calculator-api/pkg/shipping"
	"shipping-calculator-api/pkg/shipping/repository"
	"shipping-calculator-api/pkg/shipping/service"
)

// Injectors from wire.go:

func Build(db *sql.DB) shipping.Handler {
	shippingRepository := repository.Must(db)
	services := service.Must(shippingRepository)
	shippingHandler := Must(services)
	return shippingHandler
}
