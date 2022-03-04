// Функция для стандартной обработки ошибок

package handler 


type error struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Message string `json:"message"`
}


func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message) 
	c.AbortWithStatusJSON(statusCode, errorResponse{message}) // AbortWithStatusJSON - блокирует выполнение последующих обработчиков, а также записывает в 'Ответ' статус код и тело сообщения в формате JSON
}