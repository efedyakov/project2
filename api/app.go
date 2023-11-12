package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"log"

	
)

type posApi interface {	
	HealthCheck(ctx context.Context) runtime.HandlerFunc	
}

type webServicePOS interface {
	posApi
}

func runService(ctx context.Context,  server webServicePOS error {
	g, ctx := errgroup.WithContext(ctx)

	err := runHTTPServer(g, ctx,  server)
	if err != nil {
		log.Error("не удалось запустить HTTP сервис. Порт: %d, ошибка: %s", cfg.App.HTTPPort, err)
		return err
	}
	log.Info("сервис запущен по порту: %d", cfg.App.HTTPPort)

	if err = g.Wait(); err != nil {
		log.Error("", err)
	}

	return nil
}

func runHTTPServer(g *errgroup.Group, ctx context.Context, server webServicePOS) error {
	mux := runtime.NewServeMux()

	address := fmt.Sprintf(":%d", cfg.App.HTTPPort)

	httpSvc := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	_ = mux.HandlePath(http.MethodGet, "/healthz", server.HealthCheck(ctx))	
	

	g.Go(func() error {
		errLS := httpSvc.ListenAndServe()
		if errLS != nil && errLS != http.ErrServerClosed {
			log.Error("невозможно запустить http сервер: %v", errLS)
			return errLS
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		e := httpSvc.Shutdown(context.Background())
		if e == nil || !errors.Is(e, http.ErrServerClosed) {
			log.Info("http сервер остановлен")
			return nil
		}
		return e
	})

	return nil
}
