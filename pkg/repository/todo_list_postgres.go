package repository

// TodoListPostgres...
type TodoListPostgres struct {
	db *sqlx.DB
}

// NewTodoListPostgres... - метод создания списков
func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

// Create...
func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	// Begin...
	// используется для создания транзакций в метод db
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable) // Запрос для создания записи в таблицу TodoList, возвращая id нового списка
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable) // Делаем вставку в таблицу UsersList, в которой свяжем id пользователя и id нового списка
	_, err = tx.Exec(createUsersListQuery, userId, id) // Exec - метод для простого выполнения запроса, без чтения возвращаемой информации
	if err != nil {
		tx.Rollback() // Rollback - метод который откатывает все изменения БД до начала выполнения транзакции
		return 0, err // Commit - метод который применит наши изменения к БД и закончит транзакцию
	}

	return id, tx.Commit()
}

// GetAll...
func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

// GetById...
func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

// Delete...
func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

// Update...
func (r *TodoListPostgres) Update(userId, listId int, input todo.UpdateListInput) error {
	// setValues... - слайс строк
	// args... - слайс интерфейсов
	// argId... - id аргументов
	setValues := make([]string, 0) 
	args := make([]interface{}, 0)
	argId := 1

	// Проверка полей
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	// Проверка полей
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ") // setQuery - переменная в которой соеденим элементы нашего слайса в одну строку через запятую ", "

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d", // Запись запроса
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)
	
	//логируем запрос
	logrus.Debugf("updateQuery: %s", query) 
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}