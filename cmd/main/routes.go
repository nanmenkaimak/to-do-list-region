package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/handlers"
	"net/http"
)

const portNumber = ":8080"

func routes() {
	router := gin.Default()

	// возвращает 404 статус код
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	tasks := router.Group("/api/todo-list/tasks")
	{
		tasks.POST("", handlers.Repo.CreateTask)
		tasks.PUT("/:id", handlers.Repo.UpdateTask)
		tasks.DELETE("/:id", handlers.Repo.DeleteTask)
		tasks.PUT("/:id/done", handlers.Repo.Done)
		tasks.GET("", handlers.Repo.GetTasks)
	}

	_ = router.Run(portNumber)
}
