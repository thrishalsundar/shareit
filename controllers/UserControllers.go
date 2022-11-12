package controllers

import (
	"be-shareit/models"
	"be-shareit/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserRemote struct {
	Calls services.UserContracts
}

func URemoteMaker(c services.UserContracts) *UserRemote {
	return &UserRemote{
		Calls: c,
	}
}

func (app *UserRemote) UserRoutes(incomingRoutes *gin.RouterGroup) {
	routes := incomingRoutes.Group("/user")
	routes.POST("/newride", app.CreateRide)
	routes.GET("/getride", app.SpecRides)
	routes.PUT("/joinride", app.JoinRide)
	routes.PUT("/closeride", app.CloseRide)
	routes.GET("/staterides", app.StateRides)
	routes.GET("/cityrides", app.CityRides)
	routes.GET("/arearides", app.AreaRides)
	routes.GET("/history", app.ViewHistory)
	routes.PUT("/edituser", app.EditUser)
	routes.DELETE("/deleteuser", app.DeleteUser)

}
func (app *UserRemote) CreateRide(c *gin.Context) {
	userid := c.Query("userid")
	if userid == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	var newRide models.Ride
	if err := c.BindJSON(&newRide); err != nil {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.CreateRide(userid, newRide)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}
	fmt.Println(resp)
	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": nil})
	return
}

func (app *UserRemote) SpecRides(c *gin.Context) {
	userid := c.Query("user_id")

	fromAddy := c.Query("fromaddy")
	toAddy := c.Query("toaddy")

	if fromAddy == "" || toAddy == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	rides, resp := app.Calls.SpecRides(userid, fromAddy, toAddy)

	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": false, "error": resp.Err})
		return
	}

	c.JSON(200, gin.H{"msg": resp.Msg, "ok": true, "error": nil, "rides": rides})
	return
}

func (app *UserRemote) JoinRide(c *gin.Context) {
	userid := c.Query("user_id")
	rideid := c.Query("ride_id")

	if userid == "" || rideid == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.JoinRide(userid, rideid)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
	return
}

func (app *UserRemote) CloseRide(c *gin.Context) {
	userid := c.Query("user_id")
	rideid := c.Query("ride_id")

	if userid == "" || rideid == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.CloseRide(userid, rideid)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
	return
}

func (app *UserRemote) StateRides(c *gin.Context) {
	userid := c.Query("user_id")
	stateName := c.Query("state_name")

	if userid == "" || stateName == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}
	rides, resp := app.Calls.StateRides(userid, stateName)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": rides, "error": resp.Err.Error})
	return
}

func (app *UserRemote) CityRides(c *gin.Context) {
	userid := c.Query("user_id")
	cityName := c.Query("city_name")

	if userid == "" || cityName == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}
	rides, resp := app.Calls.CityRides(userid, cityName)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": rides, "error": resp.Err.Error})
	return
}

func (app *UserRemote) AreaRides(c *gin.Context) {
	userid := c.Query("user_id")
	areaName := c.Query("area_name")

	if userid == "" || areaName == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}
	rides, resp := app.Calls.AreaRides(userid, areaName)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": rides, "error": resp.Err.Error})
	return
}

func (app *UserRemote) ViewHistory(c *gin.Context) {
	userid := c.Query("user_id")
	if userid == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	rides, resp := app.Calls.ViewHistory(userid)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": rides, "error": resp.Err.Error})
	return
}

func (app *UserRemote) EditUser(c *gin.Context) {
	userid := c.Query("user_id")
	if userid == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	var changedUser models.User

	if err := c.BindJSON(&changedUser); err != nil {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.EditUser(userid, changedUser)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": resp.Err.Error})
	return
}

func (app *UserRemote) DeleteUser(c *gin.Context) {
	userid := c.Query("user_id")
	if userid == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.DeleteUser(userid)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": resp.Data, "error": resp.Err.Error})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": resp.Err.Error})
	return
}
