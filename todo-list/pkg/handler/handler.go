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
				tasks.GET("/")                //list
				tasks.POST("/", h.CreateTask) //create
				tasks.GET("/:id")             //read
				tasks.PUT("/:id")             //update
				tasks.DELETE("/:id")          //delete
			}
		}
	}

	return router
}
