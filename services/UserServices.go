package services

import (
	"be-shareit/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStruct struct {
	Users *mongo.Collection
	Rides *mongo.Collection
	ctx   context.Context
}

func UserConstruct(c context.Context, u *mongo.Collection, r *mongo.Collection) UserContracts {
	return &UserStruct{
		Users: u, Rides: r, ctx: c,
	}
}

func (app *UserStruct) CreateRide(userid string, newRide models.Ride) *models.Resp {
	var thisUser models.User
	count, err := app.Users.CountDocuments(app.ctx, bson.M{"userid": userid})
	if count <= 0 {
		return models.PResMaker("User Not Found", nil, 403, errors.New("No user Found"))
	}
	if err != nil {
		return models.PResMaker("User Count Error", nil, 500, err)
	}

	err = app.Users.FindOne(app.ctx, bson.M{"userid": userid}).Decode(&thisUser)
	if err != nil {
		return models.PResMaker("User Decode error", nil, 500, err)
	}

	newRide.ID = primitive.NewObjectID()
	newRide.RideID = newRide.ID.Hex()

	_, err = app.Rides.InsertOne(app.ctx, newRide)

	if err != nil {
		return models.PResMaker("insert error", nil, 500, err)
	}

	thisUser.History = append(thisUser.History, newRide.RideID)
	_, err = app.Users.UpdateOne(app.ctx, bson.M{"userid": userid}, bson.M{"$set": bson.M{"history": thisUser.History}})
	if err != nil {
		return models.PResMaker("Update one error", nil, 500, err)
	}

	return models.PResMaker("Successful Response", nil, 200, nil)
}

func (app *UserStruct) SpecRides(userid string, fromaddy string, toaddy string) (*[]models.Ride, *models.Resp) {
	if fromaddy == "" || toaddy == "" {
		return nil, models.PResMaker("empty string", nil, 403, errors.New("Bad request"))

	}
	var rides *[]models.Ride
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
func (app *UserStruct) JoinRide(userid string, rideid string) *models.Resp {
	var thisUser models.User
	var thisRide models.Ride
	count, err := app.Users.CountDocuments(app.ctx, bson.M{"userid": userid})
	if count <= 0 {
		return models.PResMaker("User Not Found", nil, 403, errors.New("No user Found"))
	}
	if err != nil {
		return models.PResMaker("User Count Error", nil, 500, err)
	}

	count, err = app.Rides.CountDocuments(app.ctx, bson.M{"rideid": rideid})
	if count <= 0 {
		return models.PResMaker("Ride Not Found", nil, 403, errors.New("No ride Found"))
	}
	if err != nil {
		return models.PResMaker("Ride Count Error", nil, 500, err)
	}

	err = app.Users.FindOne(app.ctx, bson.M{"userid": userid}).Decode(&thisUser)
	if err != nil {
		return models.PResMaker("User Decode error", nil, 500, err)
	}
	err = app.Rides.FindOne(app.ctx, bson.M{"rideid": rideid}).Decode(&thisRide)
	if err != nil {
		return models.PResMaker("Ride Decode error", nil, 500, err)
	}

	thisRide.Passengers = append(thisRide.Passengers, thisUser)
	thisUser.History = append(thisUser.History, thisRide.RideID)

	_, err = app.Rides.ReplaceOne(app.ctx, thisRide.RideID, thisRide)
	if err != nil {
		return models.PResMaker("replace error ride", nil, 500, err)
	}
	_, err = app.Users.ReplaceOne(app.ctx, thisUser.UserID, thisUser)
	if err != nil {
		return models.PResMaker("replace error user", nil, 500, err)
	}

	return models.PResMaker("Successful Response", nil, 200, nil)
}

func (app *UserStruct) CloseRide(userid string, rideid string) *models.Resp {
	var thisUser models.User
	var thisRide models.Ride

	count, err := app.Users.CountDocuments(app.ctx, bson.M{"userid": userid})
	if count <= 0 {
		return models.PResMaker("User Not Found", nil, 403, errors.New("No user Found"))
	}
	if err != nil {
		return models.PResMaker("User Count Error", nil, 500, err)
	}

	count, err = app.Rides.CountDocuments(app.ctx, bson.M{"rideid": rideid})
	if count <= 0 {
		return models.PResMaker("Ride Not Found", nil, 403, errors.New("No ride Found"))
	}
	if err != nil {
		return models.PResMaker("Ride Count Error", nil, 500, err)
	}

	err = app.Users.FindOne(app.ctx, bson.M{"userid": userid}).Decode(&thisUser)
	if err != nil {
		return models.PResMaker("User Decode error", nil, 500, err)
	}
	err = app.Rides.FindOne(app.ctx, bson.M{"rideid": rideid}).Decode(&thisRide)
	if err != nil {
		return models.PResMaker("Ride Decode error", nil, 500, err)
	}

	_, err = app.Rides.UpdateOne(app.ctx, bson.M{"rideid": rideid}, bson.M{"$set": bson.M{"ridestat": false}})
	if err != nil {
		return models.PResMaker("Update Document error", nil, 500, err)
	}

	return models.PResMaker("Successful Response", nil, 200, nil)
}

func (app *UserStruct) StateRides(user string, stateName string) ([]models.Ride, *models.Resp) {
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
func (app *UserStruct) CityRides(user string, cityName string) ([]models.Ride, *models.Resp) {
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
func (app *UserStruct) AreaRides(user string, areaName string) ([]models.Ride, *models.Resp) {
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

func (app *UserStruct) ViewHistory(userid string) ([]models.Ride, *models.Resp) {
	var thisUser models.User
	count, err := app.Users.CountDocuments(app.ctx, bson.M{"userid": userid})
	if count <= 0 {
		return nil, models.PResMaker("User Not Found", nil, 403, errors.New("No user Found"))
	}
	if err != nil {
		return nil, models.PResMaker("User Count Error", nil, 500, err)
	}

	err = app.Users.FindOne(app.ctx, bson.M{"userid": userid}).Decode(&thisUser)
	if err != nil {
		return nil, models.PResMaker("User Decode error", nil, 500, err)
	}

	var rides []models.Ride
	resfromdb, err := app.Rides.Find(app.ctx, bson.M{"createdby": userid})

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

func (app *UserStruct) EditUser(userid string, changedUser models.User) *models.Resp {
	var thisUser models.User
	count, err := app.Users.CountDocuments(app.ctx, bson.M{"userid": userid})
	if count <= 0 {
		return models.PResMaker("User Not Found", nil, 403, errors.New("No user Found"))
	}
	if err != nil {
		return models.PResMaker("User Count Error", nil, 500, err)
	}

	err = app.Users.FindOne(app.ctx, bson.M{"userid": userid}).Decode(&thisUser)
	if err != nil {
		return models.PResMaker("User Decode error", nil, 500, err)
	}

	_, err = app.Users.ReplaceOne(app.ctx, bson.M{"userid": userid}, changedUser)
	if err != nil {
		return models.PResMaker("Replace error", nil, 500, err)
	}

	return models.PResMaker("Successful Response", nil, 200, nil)
}

func (app *UserStruct) DeleteUser(userid string) *models.Resp {
	var thisUser models.User
	count, err := app.Users.CountDocuments(app.ctx, bson.M{"userid": userid})
	if count <= 0 {
		return models.PResMaker("User Not Found", nil, 403, errors.New("No user Found"))
	}
	if err != nil {
		return models.PResMaker("User Count Error", nil, 500, err)
	}

	err = app.Users.FindOne(app.ctx, bson.M{"userid": userid}).Decode(&thisUser)
	if err != nil {
		return models.PResMaker("User Decode error", nil, 500, err)
	}

	_, err = app.Users.DeleteOne(app.ctx, bson.M{"userid": userid})
	if err != nil {
		return models.PResMaker("Deletion error", nil, 500, err)
	}

	return models.PResMaker("Successful Response", nil, 200, nil)
}
