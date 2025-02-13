package api

import "go.uber.org/zap"

type config struct {
	addr string
	env  string
}

type app struct {
	config config
	logger *zap.SugaredLogger
}
