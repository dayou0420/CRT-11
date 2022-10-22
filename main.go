package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/crt-11/bots"
	"example.com/crt-11/controllers"
	"example.com/crt-11/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	r           *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	mongoconn := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("userdb").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	r = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World 👋!",
		})
	})

	r.POST("/callback", bots.Handler)

	basepath := r.Group("/v1")
	uc.RegisterUserRoutes(basepath)

	r.Run()
}
