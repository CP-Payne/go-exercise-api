package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CP-Payne/exercise/internal/application"
	"github.com/CP-Payne/exercise/internal/domain"
	"github.com/CP-Payne/exercise/internal/infrastructure/persistence"
	"github.com/CP-Payne/exercise/internal/interfaces/repositories"
	"github.com/CP-Payne/exercise/internal/interfaces/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewApp(cfg *config) *app {

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := persistence.NewDB(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	// Setting up routes
	router := chi.NewRouter()

	repos := repositories.NewRepositories(db)
	domainServices := domain.NewDomainServices(repos)
	applicationUseCases := application.NewApplicationUseCases(*domainServices)
	applicationHandlers := services.NewHandlers(*applicationUseCases, logger)
	applicationHandlers.RegisterRoutes(router)

	return &app{
		config: cfg,
		logger: logger,
		Router: router,
	}

}

func (app *app) run() error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.Router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infow("signal caugth", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil
}
