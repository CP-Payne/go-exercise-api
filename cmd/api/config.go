package main

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type config struct {
	addr string
	env  string
	db   dbConfig
}

type app struct {
	config *config
	logger *zap.SugaredLogger
	Router *chi.Mux
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}
