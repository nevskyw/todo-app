package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nevskyw/todo-app"
)

// Repository...
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository...
func NewRepository(db *sqlx.DB) *Repository { // db *sqlx.DB - передаем в качестве аргумента в конструктор для работы с базой данных
	// Repository... инициализируем!
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}

// Authorization...
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

// TodoList...
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

// TodoItem...
type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}
