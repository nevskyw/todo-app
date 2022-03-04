package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nevskyw/todo-app"
)

type AuthPostgres struct { // имплементирует наш интерфейс repository и работает с базой postgres
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser...
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable) // пишем запрос

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password) // QueryRow - принимает запрос и произвольный набор аргументов, которые будут подставлены в плейсхолдеры($1, $2, $3) из запроса
	if err := row.Scan(&id); err != nil {                                // с помощью метода Scan мы может записать значение id в переменную передав ее по ссылке
		return 0, err
	}

	return id, nil
}

// GetUser...
func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable) // пишем запрос
	err := r.db.Get(&user, query, username, password)                                            // Get -  передаем указатель на структуру в которую хотим записать результат выборки

	return user, err
}
