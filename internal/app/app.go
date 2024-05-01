package app

import (
	"context"
	_ "debugger-api/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gopkg.in/kothar/brotli-go.v0/enc"
	"io"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	Server          http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	a.setupHttpServer(ctx)
	return a, nil
}

func (a *App) setupHttpServer(ctx context.Context) {
	router := chi.NewMux()

	a.serviceProvider = newServiceProvider()
	a.serviceProvider.DataController()

	router.Use(middleware.URLFormat)

	compressor := middleware.NewCompressor(6)

	compressor.SetEncoder("br", func(w io.Writer, level int) io.Writer {
		params := enc.NewBrotliParams()
		params.SetQuality(level)
		return enc.NewBrotliWriter(params, w)
	})

	router.Use(compressor.Handler)

	prefix := a.serviceProvider.config.ApiPrefix

	router.Get(prefix+"/data/{key}", a.serviceProvider.dataController.GetData(ctx))
	router.Post(prefix+"/data/{key}", a.serviceProvider.dataController.PostData(ctx))
	router.Put(prefix+"/data/{key}", a.serviceProvider.dataController.UpdateData(ctx))
	router.Delete(prefix+"/data/{key}", a.serviceProvider.dataController.DeleteData(ctx))

	router.Get("/swagger/*", httpSwagger.WrapHandler) //The url pointing to API definition

	a.Server = http.Server{
		Addr:         a.serviceProvider.config.Host + ":" + a.serviceProvider.config.Port,
		Handler:      router,
		ReadTimeout:  a.serviceProvider.config.Timeout,
		WriteTimeout: a.serviceProvider.config.Timeout,
		IdleTimeout:  a.serviceProvider.config.IdleTimeout,
	}
}

func (a *App) Run() error {
	if err := a.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
