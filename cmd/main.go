package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test-task/config"
	"test-task/internal/handler"
	"test-task/internal/httpserver"
	"test-task/internal/repository"
	"test-task/internal/service"
	"test-task/pkg/database"
	"test-task/pkg/logger"
	"time"

	"go.uber.org/zap"
)

const (
	fileName = ".env"
)

func main() {

	zapLogger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	ctx, cancelContext := context.WithCancel(context.Background())

	conf, err := config.NewConfig(fileName)
	if err != nil {
		zapLogger.Error("error init config", zap.Error(err))
		return
	}

	dbPool, err := database.Initdatabase(ctx, conf, zapLogger)
	if err != nil {
		zapLogger.Error("error init database", zap.Error(err))
		return
	}

	repo := repository.NewRepositories(dbPool)
	service := service.NewService(ctx, conf, zapLogger, repo)
	handler := handler.NewHandler(service, zapLogger)
	server := httpserver.NewServer(handler.Routes())

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			zapLogger.Error("error server start", zap.Error(err))
			return
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	cancelContext()

	for i := 3; i > 0; i-- {
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	zapLogger.Info("application closed")
}
