// Логика подключения к базе, а так же храним иена таблиц в константах
package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

/*Описываем константы с названием таблиц из нашей базы в файле postgres для дальнейшего использования*/
const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

// Config...
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB... - подключение к базе данных
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	// Ping() - функция проверки
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
