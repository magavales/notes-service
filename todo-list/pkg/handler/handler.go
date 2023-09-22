package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

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
