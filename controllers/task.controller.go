package controllers

import (
	"net/http"

	"example.com/crt-11/models"
	"github.com/gin-gonic/gin"
)

type TaskService interface {
	CreateTask(*models.Task) error
}

type TaskController struct {
	TaskService TaskService
}

func New(taskservice TaskService) TaskController {
	return TaskController{
		TaskService: taskservice,
	}
}

// var taskCollection *mongo.Collection = configs.GetCollection(configs.DB, "tasks")
// var taskValidate = validator.New()

func (ts *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBind(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ts.TaskService.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
}

func (ts *TaskController) RegisterTaskRoutes(rg *gin.RouterGroup) {
	taskroute := rg.Group("/task")
	taskroute.POST("/create", ts.CreateTask)
}
