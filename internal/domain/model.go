package domain

import "time"

// Car представляет модель автомобиля
type Car struct {
	ID          int64     `json:"id"`
	RegNums     []string  `json:"regNums"`
	Mark        string    `json:"mark"`
	Model       string    `json:"model"`
	Year        int       `json:"year"`
	Owner       People    `json:"owner"`
	CreatedAt   time.Time `json:"createdAt"`
	LastUpdated time.Time `json:"lastUpdated"`
}

// People представляет модель владельца автомобиля
type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
