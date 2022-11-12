package services

import (
	"be-shareit/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GuestNeeds struct {
	ctx   context.Context
	Users *mongo.Collection
	Admin *mongo.Collection
	Rides *mongo.Collection
}

func GuestConstruct(c context.Context, u *mongo.Collection, a *mongo.Collection, r *mongo.Collection) GuestContracts {
	return &GuestNeeds{
		ctx: c, Users: u, Admin: a, Rides: r,
	}
}

func (app *GuestNeeds) Signup(newUser models.User) *models.Resp {
	ucount, err := app.Users.CountDocuments(app.ctx, bson.M{"username": newUser.Username})
	ecount, err := app.Users.CountDocuments(app.ctx, bson.M{"email": newUser.Email})
	pcount, err := app.Users.CountDocuments(app.ctx, bson.M{"phone": newUser.Phone})
	if ucount > 0 || ecount > 0 || pcount > 0 {
		return models.PResMaker("User already exist", nil, 403, errors.New("Bad request"))
	}

	newUser.ID = primitive.NewObjectID()
	newUser.UserID = newUser.ID.Hex()
	newUser.Addys = make([]models.Address, 0)
	newUser.Vehichles = make([]models.Vehichle, 0)
	newUser.History = make([]string, 0)

	_, err = app.Users.InsertOne(app.ctx, newUser)
	if err != nil {
		return models.PResMaker("Insert error", nil, 500, err)
	}

	return models.PResMaker("Successful response", nil, 200, nil)
}

func (app *GuestNeeds) Signin(uname string, passwd string) (*models.User, *models.Resp) {
	count, err := app.Users.CountDocuments(app.ctx, bson.M{"username": uname})

	if count < 1 {
		return nil, models.PResMaker("No user exists", nil, 403, errors.New("Bad request"))
	}

	var userCreds *models.User

	err = app.Users.FindOne(app.ctx, bson.M{"username": uname}).Decode(&userCreds)
	if err != nil {
		fmt.Println(err)
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	if passwd != userCreds.Password {
		return nil, models.PResMaker("Wrong Password", nil, 403, errors.New("Wrong password"))
	}

	return userCreds, models.PResMaker("Successful response", nil, 200, nil)
}

func (app *GuestNeeds) AdminSignin(uname string, passwd string) (*models.Admin, *models.Resp) {
	count, err := app.Admin.CountDocuments(app.ctx, bson.M{"username": uname})

	if count < 1 {
		return nil, models.PResMaker("No user exists", nil, 403, errors.New("Bad request"))
	}

	var adminCreds *models.Admin

	err = app.Admin.FindOne(app.ctx, bson.M{"username": uname}).Decode(&adminCreds)
	if err != nil {
		fmt.Println(err)
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	if passwd != adminCreds.Password {
		return nil, models.PResMaker("Wrong Password", nil, 403, errors.New("Wrong password"))

	}

	return adminCreds, models.PResMaker("Successful response", nil, 200, nil)

}

func (app *GuestNeeds) SpecRides(fromaddy string, toaddy string) ([]models.Ride, *models.Resp) {
	if fromaddy == "" || toaddy == "" {
		return nil, models.PResMaker("empty string", nil, 403, errors.New("Bad request"))

	}
	var rides []models.Ride
	resfromdb, err := app.Rides.Find(app.ctx, bson.M{"faddy.city": bson.M{"$regex": fromaddy}, "taddy.city": bson.M{"$regex": toaddy}})
	if err != nil {
		return nil, models.PResMaker("Find error", nil, 500, err)
	}

	err = resfromdb.All(app.ctx, &rides)
	if err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}
	defer resfromdb.Close(app.ctx)
	if err := resfromdb.Err(); err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	return rides, models.PResMaker("Successful response", nil, 200, nil)
}

func (app *GuestNeeds) AreaRides(areaName string) ([]models.Ride, *models.Resp) {
	if areaName == "" {
		return nil, models.PResMaker("empty string", nil, 403, errors.New("Bad request"))

	}
	var rides []models.Ride
	resfromdb, err := app.Rides.Find(app.ctx, bson.M{"faddy.area": bson.M{"$regex": areaName}})
	if err != nil {
		return nil, models.PResMaker("Find error", nil, 500, err)
	}

	err = resfromdb.All(app.ctx, &rides)
	if err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}
	defer resfromdb.Close(app.ctx)
	if err := resfromdb.Err(); err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	return rides, models.PResMaker("Successful response", nil, 200, nil)
}

func (app *GuestNeeds) CityRides(cityName string) ([]models.Ride, *models.Resp) {
	if cityName == "" {
		return nil, models.PResMaker("empty string", nil, 403, errors.New("Bad request"))

	}
	var rides []models.Ride
	resfromdb, err := app.Rides.Find(app.ctx, bson.M{"faddy.city": bson.M{"$regex": cityName}})
	if err != nil {
		return nil, models.PResMaker("Find error", nil, 500, err)
	}

	err = resfromdb.All(app.ctx, &rides)
	if err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}
	defer resfromdb.Close(app.ctx)
	if err := resfromdb.Err(); err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	return rides, models.PResMaker("Successful response", nil, 200, nil)
}

func (app *GuestNeeds) StateRides(stateName string) ([]models.Ride, *models.Resp) {
	if stateName == "" {
		return nil, models.PResMaker("empty string", nil, 403, errors.New("Bad request"))

	}
	var rides []models.Ride
	resfromdb, err := app.Rides.Find(app.ctx, bson.M{"faddy.state": bson.M{"$regex": stateName}})
	if err != nil {
		return nil, models.PResMaker("Find error", nil, 500, err)
	}

	err = resfromdb.All(app.ctx, &rides)
	if err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}
	defer resfromdb.Close(app.ctx)
	if err := resfromdb.Err(); err != nil {
		return nil, models.PResMaker("Decode error", nil, 500, err)
	}

	return rides, models.PResMaker("Successful response", nil, 200, nil)
}
