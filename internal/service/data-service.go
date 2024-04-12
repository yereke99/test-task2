package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"test-task/config"
	"test-task/internal/domain"
	"test-task/internal/repository"

	"go.uber.org/zap"
)

type DataService struct {
	ctx        context.Context
	zapLogger  *zap.Logger
	appConfig  *config.Config
	repository *repository.Repositories
}

func NewDataService(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repository *repository.Repositories) *DataService {

	return &DataService{
		ctx:        ctx,
		appConfig:  appConfig,
		zapLogger:  zapLogger,
		repository: repository,
	}
}

func (s *DataService) GetData() ([]*domain.Car, error) {
	car, err := s.repository.Data.GetData()
	if err != nil {
		return nil, fmt.Errorf("failed to get car data: %w", err)
	}
	return car, nil
}

func (s *DataService) DeleteData(carID int64) error {
	err := s.repository.Data.DeleteData(carID)
	if err != nil {
		return fmt.Errorf("failed to delete car data: %w", err)
	}
	return nil
}

func (s *DataService) UpdateData(car *domain.Car) error {
	err := s.repository.Data.UpdateData(car)
	if err != nil {
		return fmt.Errorf("failed to update car data: %w", err)
	}
	return nil
}

func (s *DataService) AddData(car *domain.Car) error {
	err := s.repository.Data.InsertData(car)
	if err != nil {
		return fmt.Errorf("failed to add car data: %w", err)
	}
	return nil
}

func (s *DataService) SendRequestToExternalAPI(regNums []string) error {
	// Формирование URL для запроса к внешнему API
	url := s.appConfig.Url + regNums[0]

	// Выполнение GET-запроса к внешнему API
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to send request to external API: %v", err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("external API returned non-OK status code: %d", resp.StatusCode)
	}

	// Чтение тела ответа
	var car domain.Car
	err = json.NewDecoder(resp.Body).Decode(&car)
	if err != nil {
		return fmt.Errorf("failed to decode response from external API: %v", err)
	}

	// Вывод информации о полученном автомобиле (для демонстрации)
	fmt.Println("Received car info from external API:", car)

	// Реализация отправки запроса к внешнему API
	return nil
}
