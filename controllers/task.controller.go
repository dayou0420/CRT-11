package controllers

import (
	"context"
	"net/http"
	"time"

	"example.com/crt-11/configs"
	"example.com/crt-11/models"
	"example.com/crt-11/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService interface {
	GetAll() ([]*models.Task, error)
	Create(*models.Task) error
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

func (tc *TaskController) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var tasks []models.Task
	defer cancel()

	results, err := taskCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.TaskRespose{
				Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTask models.Task
		if err = results.Decode(&singleTask); err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.TaskRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		tasks = append(tasks, singleTask)
	}

	c.JSON(http.StatusOK,
		responses.TaskRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": tasks}})

}

func (tc *TaskController) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var task models.Task
	defer cancel()

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest,
			responses.TaskRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error}})
		return
	}

	if validationErr := taskValidate.Struct(&task); validationErr != nil {
		c.JSON(http.StatusBadRequest,
			responses.TaskRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	newTask := models.Task{
		Id:   primitive.NewObjectID(),
		Name: task.Name,
	}

	result, err := taskCollection.InsertOne(ctx, newTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.TaskRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.TaskRespose{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

}

func (tc *TaskController) Routes(rg *gin.RouterGroup) {
	tr := rg.Group("/tasks")
	tr.GET("/", tc.GetAll)
	tr.POST("/", tc.Create)
}
