package app

import (
	"context"
	_ "debugger-api/docs"
	"debugger-api/internal/config"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gopkg.in/kothar/brotli-go.v0/enc"
	"io"
	"log"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	Server          http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	a.serviceProvider = newServiceProvider()

	a.setupHttpServer(ctx)
	return a, nil
}

func (a *App) setupHttpServer(ctx context.Context) {
	router := chi.NewMux()

	a.serviceProvider.DataController()

	router.Use(middleware.URLFormat)

	compressor := middleware.NewCompressor(6)

	compressor.SetEncoder("br", func(w io.Writer, level int) io.Writer {
		params := enc.NewBrotliParams()
		params.SetQuality(level)
		return enc.NewBrotliWriter(params, w)
	})

	router.Use(compressor.Handler)

	prefix := a.serviceProvider.config.HTTPServerConfig.ApiPrefix

	router.Get(prefix+"/data/{key}", a.serviceProvider.dataController.GetData(ctx))
	router.Post(prefix+"/data", a.serviceProvider.dataController.PostData(ctx))
	router.Put(prefix+"/data/{key}", a.serviceProvider.dataController.UpdateData(ctx))
	router.Delete(prefix+"/data/{key}", a.serviceProvider.dataController.DeleteData(ctx))

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	a.Server = http.Server{
		Addr:         a.serviceProvider.config.HTTPServerConfig.Host + ":" + a.serviceProvider.config.HTTPServerConfig.Port,
		Handler:      router,
		ReadTimeout:  a.serviceProvider.config.HTTPServerConfig.Timeout,
		WriteTimeout: a.serviceProvider.config.HTTPServerConfig.Timeout,
		IdleTimeout:  a.serviceProvider.config.HTTPServerConfig.IdleTimeout,
	}
}

func (a *App) Run() error {
	a.runMigrationsForPostgres()
	if err := a.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (a *App) runMigrationsForPostgres() {

	if a.serviceProvider.config.StorageConfig.Source != config.Postgres {
		return
	}

	conn := fmt.Sprintf("postgres://%s:%s@%s/%s",
		a.serviceProvider.config.StorageConfig.PostgresConfig.Username,
		a.serviceProvider.config.StorageConfig.PostgresConfig.Password,
		a.serviceProvider.config.StorageConfig.PostgresConfig.Address,
		a.serviceProvider.config.StorageConfig.PostgresConfig.Database,
	)

	dsn := flag.String("dsn", conn, "PostgreSQL data source name")

	sql, err := goose.OpenDBWithDriver("postgres", *dsn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = goose.Up(sql, "./migrations")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
