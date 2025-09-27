package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

// DB - глобальная переменная для доступа к пулу соединений БД
var DB *pgxpool.Pool

// Init инициализирует подключение к базе данных и применяет миграции
func Init() error {
	// Формируем строку подключения из переменных окружения
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Парсим конфигурацию подключения
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("error parsing connection string: %w", err)
	}

	// Настраиваем пул соединений
	config.MaxConns = 10                      // Максимальное количество соединений
	config.MinConns = 2                       // Минимальное количество соединений
	config.MaxConnLifetime = time.Hour        // Максимальное время жизни соединения
	config.MaxConnIdleTime = time.Minute * 30 // Максимальное время простоя соединения

	// Создаем пул соединений
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	// Проверяем соединение
	if err := DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// Применяем миграции
	if err := runMigrations(); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}

func checkIfMigrationsNeeded() bool {
	var tableExists bool
	err := DB.QueryRow(context.Background(),
		"SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'todos')").Scan(&tableExists)

	if err != nil {
		log.Printf("Error checking if migrations are needed: %v", err)
		return true // Если не можем проверить, лучше применить миграции
	}

	return !tableExists
}

// runMigrations применяет миграции базы данных
func runMigrations() error {
	if !checkIfMigrationsNeeded() {
		log.Println("Migrations already applied, skipping...")
		return nil
	}

	conn, err := DB.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("error acquiring connection: %w", err)
	}
	defer conn.Release()

	sqlDB := stdlib.OpenDBFromPool(DB)

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating migration driver: %w", err)
	}

	// Получаем текущую рабочую директорию
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	// Формируем правильный путь к миграциям
	migrationsPath := filepath.Join(cwd, "migrations")

	// Проверяем существование папки миграций
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		return fmt.Errorf("migrations directory does not exist: %s", migrationsPath)
	}

	// Преобразуем путь в формат URL
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %w", err)
	}

	migrationURL := "file://" + filepath.ToSlash(absPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationURL,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	log.Println("Database migrations applied successfully")
	return nil
}

// Close закрывает все соединения с базой данных
func Close() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
