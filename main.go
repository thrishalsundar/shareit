// guest->seller->admin->user
package main

import (
	"be-shareit/controllers"
	"be-shareit/database"
	"be-shareit/services"
	"context"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server *gin.Engine
	port   string
	client *mongo.Client
	ctx    context.Context
	Users  *mongo.Collection
	Rides  *mongo.Collection
	Admins *mongo.Collection

	GServes    services.GuestContracts
	GControlla *controllers.GuestRemote
	AServes    services.AdminContracts
	AControlla *controllers.AdminRemote
	UServes    services.UserContracts
	UControlla *controllers.UserRemote
)

func init() {
	godotenv.Load()
	port = os.Getenv("PORT")
	if port == "" {
		panic("Port ah kaanum")
	}

	client, ctx = database.DbSetup()
	if client == nil || ctx == nil {
		panic("client not created error")
	}
	Users = client.Database("shareit").Collection("users")
	Rides = client.Database("shareit").Collection("rides")
	Admins = client.Database("shareit").Collection("admins")

	GServes = services.GuestConstruct(ctx, Users, Admins, Rides)
	GControlla = controllers.GRemoteMaker(GServes)

	AServes = services.AdminConstruct(ctx, Users, Admins, Rides)
	AControlla = controllers.ARemoteMaker(AServes)

	UServes = services.UserConstruct(ctx, Users, Rides)
	UControlla = controllers.URemoteMaker(UServes)

	server = gin.Default()
	server.Use(cors.Default())

}

func main() {
	defer client.Disconnect(ctx)

	basePath := server.Group("/apis")

	GControlla.GuestRoutes(basePath)

	//authmiddelware

	//authedRoutes

	UControlla.UserRoutes(basePath)
	AControlla.AdminRoutes(basePath)

	server.Run(":" + port)

	fmt.Println("vim-go")
}
