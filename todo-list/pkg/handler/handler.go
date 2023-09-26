package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todo-list/docs"
	"todo-list/pkg/database"
)

type Handler struct {
	db *database.Database
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			tasks := v1.Group("/tasks")
			{
				tasks.GET("", h.getTasks)              //list
				tasks.POST("", h.createTask)           //create
				tasks.GET("/:id", h.getTaskByID)       //read
				tasks.PUT("/:id", h.updateTaskByID)    //update
				tasks.DELETE("/:id", h.deleteTaskByID) //delete
			}
		}
	}

	return router
}
