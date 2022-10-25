package controllers

import (
	"net/http"

	"example.com/crt-11/configs"
	"example.com/crt-11/models"
	"example.com/crt-11/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

var taskCollection *mongo.Collection = configs.GetCollection(configs.DB, "tasks")
var taskValidate = validator.New()

func (tc *TaskController) Create(ctx *gin.Context) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var task models.Task
	// defer cancel()

	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest,
			responses.CityRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error}})
		return
	}

	if validationErr := taskValidate.Struct(&task); validationErr != nil {
		ctx.JSON(http.StatusBadRequest,
			responses.CityRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	newTask := models.City{
		Id:   primitive.NewObjectID(),
		Name: task.Name,
	}

	result, err := taskCollection.InsertOne(ctx, newTask)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			responses.CityRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusCreated, responses.CityRespose{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

}

func (tc *TaskController) RegisterTaskRoutes(rg *gin.RouterGroup) {
	taskroute := rg.Group("/task")
	taskroute.POST("/create", tc.Create)
}
