package repository

import (
	"database/sql"
	"test-task/internal/domain"
)

type IData interface {
	GetData() ([]*domain.Car, error)
	DeleteData(carID int64) error
	UpdateData(car *domain.Car) error
	InsertData(car *domain.Car) error
}

type Repositories struct {
	Data IData
}

func NewRepositories(db *sql.DB) *Repositories {

	return &Repositories{
		Data: NewDataRepository(db),
	}
}
