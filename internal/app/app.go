package app

import (
	"context"
	_ "debugger-api/docs"
	"debugger-api/internal/config"
	"debugger-api/internal/middleware"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	Server          http.Server
}

func NewApp(ctx context.Context, cfg config.Config) (*App, error) {
	a := &App{}
	a.serviceProvider = newServiceProvider(cfg)

	err := a.setup(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) setup(ctx context.Context) error {

	funcs := []func(context.Context) error{
		a.runMigrationsForPostgres,
		a.setupHttpServer,
	}

	for _, f := range funcs {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) setupHttpServer(ctx context.Context) error {

	// TODO return compression
	//compressor := middleware.NewCompressor(6)
	//
	//compressor.SetEncoder("br", func(w io.Writer, level int) io.Writer {
	//	params := enc.NewBrotliParams()
	//	params.SetQuality(level)
	//	return enc.NewBrotliWriter(params, w)
	//})

	router := http.NewServeMux()

	mockRouter := http.NewServeMux()

	mockRouter.HandleFunc("GET /{id}/content", a.serviceProvider.DataController().GetMockContent(ctx))
	mockRouter.HandleFunc("GET /{id}", a.serviceProvider.DataController().GetMock(ctx))
	mockRouter.HandleFunc("DELETE /{id}", a.serviceProvider.DataController().DeleteMock(ctx))
	mockRouter.HandleFunc("PATCH /{id}", a.serviceProvider.DataController().UpdateMock(ctx))

	apiRouter := http.NewServeMux()

	mockPrefix := "/mock"
	apiRouter.Handle(mockPrefix+"/",
		http.StripPrefix(mockPrefix, mockRouter))
	apiRouter.HandleFunc("POST /mock", a.serviceProvider.DataController().PostMock(ctx))

	prefix := a.serviceProvider.config.HTTPServerConfig.ApiPrefix
	router.Handle(prefix+"/",
		http.StripPrefix(prefix, apiRouter))

	router.Handle("/swagger/*", httpSwagger.WrapHandler)

	middlewareStack := middleware.CreateStack(
		a.serviceProvider.RequestIdMiddlewareProvider().GetRequestIdMiddleware,
		a.serviceProvider.FullRequestUrlMiddlewareProvider().GetFullRequestUrlMiddleware,
		//Not working now
		a.serviceProvider.CompressorMiddlewareProvider().GetCompressorMiddleware,
	)

	a.Server = http.Server{
		Addr:         a.serviceProvider.config.HTTPServerConfig.Host + ":" + a.serviceProvider.config.HTTPServerConfig.Port,
		Handler:      middlewareStack(router),
		ReadTimeout:  a.serviceProvider.config.HTTPServerConfig.Timeout,
		WriteTimeout: a.serviceProvider.config.HTTPServerConfig.Timeout,
		IdleTimeout:  a.serviceProvider.config.HTTPServerConfig.IdleTimeout,
	}

	return nil
}

func (a *App) Run() error {
	return a.Server.ListenAndServe()
}

func (a *App) Stop() error {
	return a.Server.Shutdown(context.Background())
}

func (a *App) runMigrationsForPostgres(_ context.Context) error {

	if a.serviceProvider.config.StorageConfig.Source != config.Postgres {
		return nil
	}

	conn := fmt.Sprintf("postgres://%s:%s@%s/%s",
		a.serviceProvider.config.StorageConfig.PostgresConfig.Username,
		a.serviceProvider.config.StorageConfig.PostgresConfig.Password,
		a.serviceProvider.config.StorageConfig.PostgresConfig.Address,
		a.serviceProvider.config.StorageConfig.PostgresConfig.Database,
	)

	dsn := flag.String("dsn", conn, "PostgreSQL mock source name")

	sql, err := goose.OpenDBWithDriver("postgres", *dsn)
	if err != nil {
		return err
	}

	err = goose.Up(sql, "./migrations")
	if err != nil {
		return err
	}

	return nil
}
