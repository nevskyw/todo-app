package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nevskyw/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service // внедрение зависимостей | указатель на service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// Инициализация роутера
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // объявляем наши методы сгрупировав их по маршрутам

	auth := router.Group("/auth") // метод /auth - для регистрации и авторизации
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity) // будет использоваться /api для рыботы со СПИСКАМИ и их ЗАДАЧАМИ
	{
		lists := api.Group("/lists") // работа со списками - /lists
		{
			lists.POST("/", h.createList)      // создание
			lists.GET("/", h.getAllLists)      // получение всех списков
			lists.GET("/:id", h.getListById)   // получение списка по ID
			lists.PUT("/:id", h.updateList)    // редактирование списка
			lists.DELETE("/:id", h.deleteList) // удаление

			items := lists.Group(":id/items") // Задачи списка  - ":id/items"
			{
				items.POST("/", h.createItem) // создание
				items.GET("/", h.getAllItems) // удаление
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router // возвращение роутера
}
