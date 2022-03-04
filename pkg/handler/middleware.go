/*Прослойка - которая будет парсить JWT токены из запроса и предоставлять доступ к нашей группе эндпоинтов /api */

package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

// userIdentity... - обработчик
func (h *Handler) userIdentity(c *gin.Context) {
	// получаем значения из authorizationHeader...
	// валидируем что он не пустой
	// при ошибке возвращаем статус код 401
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	// stringSplit... - вызываем функцию и указываем разделить нашу строку по пробелам
	// возвращаем массив длинною в 2 элемента
	// при ошибке возвращаем статус код 401
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	// ParseToken...
	/* если все успешно, то записываем значение id в контекст, это делается для того, чтобы иметь доступ к id пользователя,
	который делает запрос. В последующих обработчиках, которые вызываются после данной прослойки*/
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

// getUserId...
// Функция где мы обрабатываем ошибки
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
