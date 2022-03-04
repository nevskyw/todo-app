package todo

// User struct...
type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`     // binding:"required - тег, который валидирует наличие данных полей в теле запроса
	Username string `json:"username" binding:"required"` // binding:"required - тег, который валидирует наличие данных полей в теле запроса
	Password string `json:"password" binding:"required"` // binding:"required - тег, который валидирует наличие данных полей в теле запроса
}
