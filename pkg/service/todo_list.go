package service

import (
	"github.com/nevskyw/todo-app"
	"github.com/nevskyw/todo-app/pkg/repository"
)
// TodoListService... передаем в качестве поля нашей структуры и будем передовать в конструкторе
type TodoListService struct {
	repo repository.TodoList
}

// NewTodoListService...
func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Create...
func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

// GetAll...
func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

// GetById...
func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

// Delete...
func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

// Update...
func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
