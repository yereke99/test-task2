package database

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"test-task/config"
	"test-task/internal/domain"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Initdatabase(ctx context.Context, databaseConfig *config.Config, zapLogger *zap.Logger) (*sql.DB, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Database,
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	zapLogger.Info("Connected to PostgreSQL database", zap.Any("database", databaseConfig.Database))
	return db, nil
}

func Migrate(ctx context.Context, db *sql.DB, zapLogger *zap.Logger) error {

	query := `SELECT * FROM car`
	_, err := db.Query(query)
	if err == nil {
		return domain.ErrExistsTable
	}

	// Read migration files
	upSQL, err := ioutil.ReadFile("path/to/0001_init_schema.up.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration up file: %v", err)
	}

	// Apply migration
	if _, err := db.ExecContext(ctx, string(upSQL)); err != nil {
		return fmt.Errorf("failed to apply migration: %v", err)
	}

	zapLogger.Info("applied migration")

	return nil
}
