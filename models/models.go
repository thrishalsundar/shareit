package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   string             `json:"user_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    string             `json:"user_id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Addys     []Address          `bson:"address" json:"address"`
	Vehichles []Vehichle         `bson:"vehichle" json:"vehichle"`
	History   []string           `json:"history"`
}

type Address struct {
	ID           primitive.ObjectID `bson:"_id"`
	AddressID    string             `json:"address_id"`
	AddressTitle string             `json:"address_title"`
	DoorNo       string             `json:"door_no"`
	Street       string             `json:"street"`
	Area         string             `json:"area"`
	City         string             `json:"city"`
	PinCode      string             `json:"pincode"`
	State        string             `json:"state"`
}

type Ride struct {
	ID            primitive.ObjectID `bson:"_id"`
	RideID        string             `json:"ride_id"`
	CreatedBy     string             `json:"createdby"`
	OverallAmount uint64             `json:"totamount"`
	HeadsNeeded   uint32             `json:"heads_needed"`
	PerHead       uint32             `json:"perhead"`
	FAddy         Address            `bson:"from_address" json:"from_address"`
	TAddy         Address            `bson:"to_address" json:"to_address"`
	SeatsLeft     uint32             `json:"seatsleft"`
	RideStat      bool               `json:"ride_stat"`
	RideAt        time.Time          `json:"ride_at"`
	Passengers    []User             `bson:"passengers" json:"passengers"`
	Tags          []string           `json:"tags"`
}

type Vehichle struct {
	ID           primitive.ObjectID `bson:"_id"`
	VehichleID   string             `json:"vehichle_id"`
	VehichleName string             `json:"vehichle_name"`
	ImageLinks   []string           `json:"image_links"`
	VehichleNo   string             `json:"vehichleno"`
	Color        string             `json:"color"`
}
