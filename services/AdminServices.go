package services

import (
	"be-shareit/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminStruct struct {
	ctx    context.Context
	Users  *mongo.Collection
	Admins *mongo.Collection
	Rides  *mongo.Collection
}

func AdminConstruct(c context.Context, u *mongo.Collection, a *mongo.Collection, r *mongo.Collection) AdminContracts {
	return &AdminStruct{
		ctx: c, Users: u, Admins: a, Rides: r,
	}
}

func (app *AdminStruct) ShowRides() ([]models.Ride, *models.Resp) {

	var allRides []models.Ride
	cursor, err := app.Rides.Find(app.ctx, bson.M{})
	if err != nil {
		return nil, models.PResMaker("Find error", nil, 500, err)
	}

	err = cursor.All(app.ctx, &allRides)
	if err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}
	defer cursor.Close(app.ctx)
	if err := cursor.Err(); err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	return allRides, models.PResMaker("Successfull Response", allRides, 200, nil)
}

func (app *AdminStruct) ShowUsers() ([]models.User, *models.Resp) {
	var allUsers []models.User
	cursor, err := app.Users.Find(app.ctx, bson.M{})
	if err != nil {
		return nil, models.PResMaker("Find error", nil, 500, err)
	}

	err = cursor.All(app.ctx, &allUsers)
	if err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}
	defer cursor.Close(app.ctx)
	if err := cursor.Err(); err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	return allUsers, models.PResMaker("Successfull Response", allUsers, 200, nil)

}

func (app *AdminStruct) EditUser(editedUser models.User) *models.Resp {

	_, err := app.Users.UpdateOne(app.ctx, bson.M{"userid": editedUser.ID}, editedUser)
	if err != nil {
		return models.PResMaker("Update Error", nil, 500, err)
	}

	return models.PResMaker("Successful Response", nil, 200, nil)
}

func (app *AdminStruct) DeleteUsers(users []string) *models.Resp {

	var queries []mongo.WriteModel
	for _, i := range users {
		queries = append(queries, mongo.NewDeleteOneModel().SetFilter(bson.M{"userid": i}))
	}

	res, err := app.Users.BulkWrite(app.ctx, queries)
	if err != nil {
		return models.PResMaker("Bulkwrite error", nil, 500, err)
	}
	fmt.Println(res)

	return models.PResMaker("Successful Response", res, 200, err)
}

func (app *AdminStruct) DeleteRides(rides []string) *models.Resp {
	var queries []mongo.WriteModel
	for _, i := range rides {
		queries = append(queries, mongo.NewDeleteOneModel().SetFilter(bson.M{"userid": i}))
	}

	res, err := app.Rides.BulkWrite(app.ctx, queries)
	if err != nil {
		return models.PResMaker("Bulkwrite error", nil, 500, err)
	}
	fmt.Println(res)

	return models.PResMaker("Successful Response", res, 200, err)
}
