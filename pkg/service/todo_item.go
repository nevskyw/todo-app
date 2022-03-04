package service

import (
	"github.com/nevskyw/todo-app"
	"github.com/nevskyw/todo-app/pkg/repository"
)

// TodoItemService...
type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

// NewTodoItemService...
func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

// Create...
func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// список не существует или не принадлежит пользователю
		return 0, err
	}

	return s.repo.Create(listId, item)
}

// GetAll...
func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

// GetById...
func (s *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

// Delete...
func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

// Update...
func (s *TodoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
