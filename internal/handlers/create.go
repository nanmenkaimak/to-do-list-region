package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"net/http"
	"time"
)

// @Summary Create New Task
// @Description Создает новые задачи
// @ID create-task
// @Accept json
// @Param input body models.Task true "task values"
// @Success 204
// @Failure 400,404 {object} error
// @Router /api/todo-list/tasks [post]
// CreateTask создает новые задачи
func (m *Repository) CreateTask(ctx *gin.Context) {
	var newTask models.Task

	// обработка JSON
	if err := ctx.BindJSON(&newTask); err != nil {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// проверяет что title и activeAt не пустые
	if newTask.Title == "" || newTask.ActiveAt == "" {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("missing values").Error()})
		return
	}

	// преобразует activeAt в Time
	_, err := time.Parse("2006-01-02", newTask.ActiveAt)
	if err != nil {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// записывает данные в базу
	err = m.DB.CreateTask(newTask)
	if err != nil {
		// возвращает 404 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// возвращает 204 статус код
	ctx.Status(http.StatusNoContent)
}
