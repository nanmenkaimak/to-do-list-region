package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// Done помечает задачу выполненной
func (m *Repository) Done(ctx *gin.Context) {
	// берет ID из URL и преобразует в primitive.ObjectID
	taskID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// обновняет статус задачи
	err = m.DB.UpdateStatus(taskID)
	if err != nil {
		// возвращает 404 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// возвращает 204 статус код
	ctx.Status(http.StatusNoContent)
}
