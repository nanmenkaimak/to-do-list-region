package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// @Summary Update Task
// @Description Обновляет уже существующий задачи
// @ID update-task
// @Accept json
// @Param input body models.Task true "task values"
// @Param id   path string  true  "Task ID"
// @Success 204
// @Failure 400,404 {object} error
// @Router /api/todo-list/tasks/{id} [put]
// UpdateTask обновляет уже существующий задачи
func (m *Repository) UpdateTask(ctx *gin.Context) {
	var updatedTask models.Task

	// обработка JSON
	if err := ctx.BindJSON(&updatedTask); err != nil {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// берет ID из URL и преобразует в primitive.ObjectID
	taskID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask.ID = taskID

	// обновляет задачу
	err = m.DB.UpdateTask(updatedTask)
	if err != nil {
		// возвращает 404 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// возвращает 204 статус код
	ctx.Status(http.StatusNoContent)
}
