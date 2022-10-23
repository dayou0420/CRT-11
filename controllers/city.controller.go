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

var cityCollection *mongo.Collection = configs.GetCollection(configs.DB, "cities")
var cityValidate = validator.New()

func GetAllCities() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var cities []models.City
		defer cancel()

		results, err := cityCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.CityRespose{
					Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCity models.City
			if err = results.Decode(&singleCity); err != nil {
				c.JSON(http.StatusInternalServerError,
					responses.CityRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			cities = append(cities, singleCity)
		}

		c.JSON(http.StatusOK,
			responses.CityRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": cities}})
	}
}

func GetCity() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancal := context.WithTimeout(context.Background(), 10*time.Second)
		cityId := c.Param("cityId")
		var city models.City
		defer cancal()

		objId, _ := primitive.ObjectIDFromHex(cityId)

		err := cityCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&city)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.CityRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.CityRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": city}})
	}
}

func CreateCity() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var city models.City
		defer cancel()

		if err := c.BindJSON(&city); err != nil {
			c.JSON(http.StatusBadRequest,
				responses.CityRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := cityValidate.Struct(&city); validationErr != nil {
			c.JSON(http.StatusBadRequest,
				responses.CityRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newCity := models.City{
			Id:        primitive.NewObjectID(),
			Name:      city.Name,
			Latitude:  city.Latitude,
			Longitude: city.Longitude,
		}

		result, err := cityCollection.InsertOne(ctx, newCity)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.CityRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.CityRespose{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func UpdateCity() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cityId := c.Param("cityId")
		var city models.City
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cityId)

		if err := c.BindJSON(&city); err != nil {
			c.JSON(http.StatusBadRequest,
				responses.CityRespose{
					Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validateErr := cityValidate.Struct(&city); validateErr != nil {
			c.JSON(http.StatusBadRequest,
				responses.CityRespose{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateErr.Error()}})
			return
		}

		update := bson.M{"name": city.Name, "latitude": city.Latitude, "longitude": city.Longitude}
		result, err := cityCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.CityRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updateCity models.City
		if result.MatchedCount == 1 {
			err := cityCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateCity)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					responses.CityRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.CityRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateCity}})
	}
}

func DeleteCity() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cityId := c.Param("cityId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cityId)

		result, err := cityCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.CityRespose{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.CityRespose{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "City with specified Id not found."}})
			return
		}

		c.JSON(http.StatusOK, responses.CityRespose{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "City successfully deleted."}})
	}
}
