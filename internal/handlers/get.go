package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"net/http"
	"time"
)

// GetTasks показывает список задач по статусу
func (m *Repository) GetTasks(ctx *gin.Context) {
	// берет значение из query, если статус = done, тогда statusBool изменяется на true
	status := ctx.Query("status")
	statusBool := false
	if status == "done" {
		statusBool = true
	}

	// возвращает все задачи у которых activeAt <= текущего дня и по статусу
	allTasks, err := m.DB.GetAllTasks(statusBool)
	if err != nil {
		// возвращает 400 статус код и ошибку
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// если за это день нет задач, тогда возвращается пустой массив []
	if allTasks == nil {
		allTasks = []models.Task{}
	}

	// проверяет что день попадает в выходные дни
	for i := range allTasks {
		// преобразует activeAt в Time
		date, err := time.Parse("2006-01-02", allTasks[i].ActiveAt)
		if err != nil {
			// возвращает 400 статус код и ошибку
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// проверияет что день попадает в субботу или в воскресенье
		if int(date.Weekday()) == 6 || int(date.Weekday()) == 0 {
			allTasks[i].Title = fmt.Sprintf("ВЫХОДНОЙ - %s", allTasks[i].Title)
		}
	}
	// возвращает 200 статус код
	ctx.JSON(http.StatusOK, allTasks)
}
