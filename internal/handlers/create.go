package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"net/http"
	"time"
)

func (m *Repository) CreateTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newTask.Title == "" || newTask.ActiveAt == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("missing values").Error()})
		return
	}

	_, err := time.Parse("2006-01-02", newTask.ActiveAt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = m.DB.CreateTask(newTask)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
