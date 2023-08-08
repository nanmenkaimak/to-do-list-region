package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// @Summary Delete Task
// @Description Удаляет задачу
// @ID delete-task
// @Param id   path string  true  "Task ID"
// @Success 204
// @Failure 400,404 {object} error
// @Router /api/todo-list/tasks/{id} [delete]
// DeleteTask удаляет задачу
func (m *Repository) DeleteTask(ctx *gin.Context) {
	// берет ID из URL и преобразует в primitive.ObjectID
	taskID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// удаляет задачу из базы данных
	err = m.DB.DeleteTask(taskID)
	if err != nil {
		// возвращает 404 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// возвращает 204 статус код
	ctx.Status(http.StatusNoContent)
}
