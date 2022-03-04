// Обработчик регистрации и авторизации
package handler 

// signUp ... 
// регистрация
func (h *Handler) signUp(c *gin.Context) { 
	var input todo.User // - будем записывать данные из JSON для пользователей

	if err := c.BindJSON(&input); err != nil { // парсим JSON
		newErrorResponse(c, http.StatusBadRequest, "invalid input body") //newErrorResponse - функция для создания ответа с ошибкой
		return
	}

	id, err := h.services.Authorization.CreateUser(input) // вызываем CreateUser в которую мы передадим нашу структуру пользователя
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // если будет ошибка, то вызовем функцию для записи ответа с ошибкой и статус код 500
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{ // если все ОК, то записываем в ответ статус код 200 и тело JSON со значением id пользователя
		"id": id,
	})
}

// signInInput... получение логина и пароля
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


// signIn ... 
// авторизация
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput // для получения логина и пороля в запросе

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password) // GenerateToken - замена при авторизации на токен
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}