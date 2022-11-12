package services

import "be-shareit/models"

type UserContracts interface {
	CreateRide(user string, ride models.Ride) *models.Resp
	SpecRides(user string, fromAddy string, toAddy string) (*[]models.Ride, *models.Resp)
	JoinRide(user string, rideid string) *models.Resp
	CloseRide(user string, rideid string) *models.Resp
	StateRides(user string, stateName string) ([]models.Ride, *models.Resp)
	CityRides(user string, cityName string) ([]models.Ride, *models.Resp)
	AreaRides(user string, areaName string) ([]models.Ride, *models.Resp)
	ViewHistory(user string) ([]models.Ride, *models.Resp)
	EditUser(user string, changedUser models.User) *models.Resp
	DeleteUser(user string) *models.Resp
}

type AdminContracts interface {
	ShowRides() ([]models.Ride, *models.Resp)
	ShowUsers() ([]models.User, *models.Resp)
	EditUser(models.User) *models.Resp
	DeleteUsers([]string) *models.Resp
	DeleteRides([]string) *models.Resp
}

type GuestContracts interface {
	Signup(models.User) *models.Resp
	Signin(string, string) (*models.User, *models.Resp)
	AdminSignin(string, string) (*models.Admin, *models.Resp)
	SpecRides(string, string) ([]models.Ride, *models.Resp)
	AreaRides(string) ([]models.Ride, *models.Resp)
	CityRides(string) ([]models.Ride, *models.Resp)
	StateRides(string) ([]models.Ride, *models.Resp)
}
