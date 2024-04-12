package repository

import (
	"database/sql"
	"fmt"
	"test-task/internal/domain"
	"time"
)

type Data struct {
	db *sql.DB
}

func NewDataRepository(db *sql.DB) *Data {

	return &Data{
		db: db,
	}
}

func (r *Data) GetData() ([]*domain.Car, error) {
	var cars []*domain.Car

	rows, err := r.db.Query("SELECT id, reg_nums, mark, model, year, owner_name, owner_surname, owner_patronymic, created_at, last_updated FROM cars")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var car domain.Car
		err := rows.Scan(&car.ID, &car.RegNums, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic, &car.CreatedAt, &car.LastUpdated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		cars = append(cars, &car)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return cars, nil
}

func (r *Data) DeleteData(carID int64) error {
	result, err := r.db.Exec("DELETE FROM cars WHERE id = $1", carID)
	if err != nil {
		return fmt.Errorf("failed to delete car: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no car found with ID: %d", carID)
	}

	return nil
}

func (r *Data) UpdateData(car *domain.Car) error {
	_, err := r.db.Exec("UPDATE cars SET reg_nums = $1, mark = $2, model = $3, year = $4, owner_name = $5, owner_surname = $6, owner_patronymic = $7, last_updated = $8 WHERE id = $9",
		car.RegNums, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic, time.Now(), car.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update car: %w", err)
	}
	return nil
}

func (r *Data) InsertData(car *domain.Car) error {
	_, err := r.db.Exec("INSERT INTO cars (reg_nums, mark, model, year, owner_name, owner_surname, owner_patronymic, created_at, last_updated) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		car.RegNums, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic, time.Now(), time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert car: %w", err)
	}
	return nil
}
