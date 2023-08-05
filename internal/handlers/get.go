package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"net/http"
	"time"
)

func (m *Repository) GetTasks(ctx *gin.Context) {
	status := ctx.Query("status")
	statusBool := false
	if status == "done" {
		statusBool = true
	}

	allTasks, err := m.DB.GetAllTask(statusBool)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if allTasks == nil {
		allTasks = []models.Task{}
	}

	for i := range allTasks {
		date, err := time.Parse("2006-01-02", allTasks[i].ActiveAt)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if int(date.Weekday()) == 6 || int(date.Weekday()) == 0 {
			allTasks[i].Title = fmt.Sprintf("ВЫХОДНОЙ - %s", allTasks[i].Title)
		}
	}
	ctx.JSON(http.StatusOK, allTasks)
}
