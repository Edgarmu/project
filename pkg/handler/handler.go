package handler

import (
	"todo/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(servises *service.Service) *Handler {
	return &Handler{services: servises}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}
	api := router.Group("api")
	{
		list := api.Group("/list")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllList)
			list.GET("/:id", h.getListById)
			list.PUT("/:id", h.updateList)
			list.DELETE("/id", h.deleteList)

			item := list.Group(":id/items")
			{
				item.POST("/", h.createItem)
				item.GET("/", h.getAllItem)
				item.GET("/:item_id", h.getItemById)
				item.PUT("/:item_id", h.updateItem)
				item.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
