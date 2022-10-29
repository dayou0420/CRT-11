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
	Get(*string) (*models.Task, error)
	Create(*models.Task) error
	Update(*models.Task) error
	Delete(*string) error
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

func (tc *TaskController) Get(c *gin.Context) {
	ctx, cancal := context.WithTimeout(context.Background(), 10*time.Second)
	taskId := c.Param("id")
	var task models.Task
	defer cancal()

	objId, _ := primitive.ObjectIDFromHex(taskId)

	err := taskCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.TaskRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK,
		responses.TaskRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": task}})
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
		Date: task.Date,
		Bill: task.Bill,
		Account: models.Account{
			Name:    task.Account.Name,
			State:   task.Account.State,
			City:    task.Account.City,
			Pincode: task.Account.Pincode,
		},
	}

	result, err := taskCollection.InsertOne(ctx, newTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.TaskRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated,
		responses.TaskRespose{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}

func (tc *TaskController) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	taskId := c.Param("id")
	var task models.Task
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(taskId)

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest,
			responses.TaskRespose{
				Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validateErr := taskValidate.Struct(&task); validateErr != nil {
		c.JSON(http.StatusBadRequest,
			responses.TaskRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateErr.Error()}})
		return
	}

	update := bson.M{"name": task.Name, "date": task.Date, "bill": task.Bill, "account": models.Account{}}
	result, err := taskCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.TaskRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	var updateTask models.Task
	if result.MatchedCount == 1 {
		err := taskCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateTask)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.TaskRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
	}

	c.JSON(http.StatusOK, responses.CityRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateTask}})
}

func (tc *TaskController) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	taskId := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(taskId)

	result, err := taskCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.TaskRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound,
			responses.TaskRespose{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Task with specified id not found."}})
		return
	}

	c.JSON(http.StatusOK, responses.TaskRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Task successfully deleted."}})
}

func (tc *TaskController) TaskRoute(rg *gin.RouterGroup) {
	r := rg.Group("/tasks")
	r.GET("/", tc.GetAll)
	r.GET("/:id", tc.Get)
	r.POST("/", tc.Create)
	r.PATCH("/:id", tc.Update)
	r.DELETE("/:id", tc.Delete)
}
