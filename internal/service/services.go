package service

import (
	"context"
	"test-task/config"
	"test-task/internal/domain"
	"test-task/internal/repository"

	"go.uber.org/zap"
)

type IService interface {
	SendRequestToExternalAPI(regNums []string) error
	GetData() ([]*domain.Car, error)
	DeleteData(carID int64) error
	UpdateData(car *domain.Car) error
	AddData(car *domain.Car) error
}

type Service struct {
	Service IService
}

func NewService(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repository *repository.Repositories) *Service {

	return &Service{}
}
