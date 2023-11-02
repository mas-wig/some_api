package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mas-wig/post-api-1/api"
	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/routes"
	"github.com/mas-wig/post-api-1/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client
	redisClient *redis.Client

	userServices      services.UserService
	userHandler       api.UserHandler
	userRoutesHandler routes.UserRoutesHandler

	authCollection    *mongo.Collection
	authServices      services.AuthService
	authHandler       api.AuthHandler
	authRoutesHandler routes.AuthRoutesHandler
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not find env file in root dir")
	}
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI(config.MongoDBURL)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal("could not connect to database")
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Success make connection to mongodb...")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisURI,
		Password: config.RedisPassword,
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		panic(err)
	}
	err = redisClient.Set("test", "Welcome to hell", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis client connected successfully.....")
	server = gin.Default()
}

func main() {
	config, _ := config.LoadConfig(".")
	value, err := redisClient.Get("test").Result()

	if err == redis.Nil {
		fmt.Println("key : test does not exist")
	} else if err != nil {
		panic(err)
	}

	server.Use(cors.New(cors.Config{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		MaxAge:           0,
	}))

	router := server.Group("/api/")

	authRoutesHandler.AuthRouters(router, userServices)
	userRoutesHandler.UserRouters(router, userServices)

	router.GET("healtchecker", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": &value})
	})

	defer mongoClient.Disconnect(ctx)
	log.Fatal(server.Run(":" + config.PORT))
}
