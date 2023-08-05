package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/handlers"
	"net/http"
)

const portNumber = ":8080"

func routes() {
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	tasks := router.Group("/api/todo-list/tasks")
	{
		tasks.POST("", handlers.Repo.CreateTask)
		tasks.PUT("/:id", handlers.Repo.UpdateTask)
	}

	_ = router.Run(portNumber)
}
