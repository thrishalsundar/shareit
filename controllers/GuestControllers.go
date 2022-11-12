package controllers

import (
	"be-shareit/models"
	"be-shareit/services"

	"github.com/gin-gonic/gin"
)

type GuestRemote struct {
	Calls services.GuestContracts
}

func GRemoteMaker(c services.GuestContracts) *GuestRemote {
	return &GuestRemote{
		Calls: c,
	}
}

func (app *GuestRemote) GuestRoutes(incomingRoutes *gin.RouterGroup) {
	routes := incomingRoutes.Group("/guest")
	routes.POST("/signup", app.Signup)
	routes.POST("/signin", app.Signin)
	routes.GET("/routerides", app.SpecRides)
	routes.GET("/arearides", app.AreaRides)
	routes.GET("/cityrides", app.CityRides)
	routes.GET("/staterides", app.StateRides)
}

func (app *GuestRemote) Signup(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(403, gin.H{"msg": "Parse error", "ok": false, "error": err})
		return
	}

	resp := app.Calls.Signup(newUser)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": false, "error": resp.Err})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": true, "error": nil})
	return
}

func (app *GuestRemote) Signin(c *gin.Context) {
	type LoginStruct struct {
		Uname    string `json:"uname"`
		Password string `json:"password"`
	}
	var loginReq LoginStruct

	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(403, gin.H{"msg": "Parsing Error", "err": err, "ok": false})
		return
	}

	if loginReq.Password == "" || loginReq.Uname == "" {
		c.JSON(403, gin.H{"msg": "Parsing Error", "err": "parse", "ok": false})
		return
	}

	authedUser, resp := app.Calls.Signin(loginReq.Uname, loginReq.Password)
	if resp.Msg == "Wrong Password" {
		c.JSON(403, gin.H{"msg": "Wrong Password", "ok": false})
		return
	}
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": false, "err": resp.Err})
		return
	}

	c.JSON(200, gin.H{"msg": "Authenticated Successfully", "ok": true, "user": authedUser})
	return
}

func (app *GuestRemote) SpecRides(c *gin.Context) {
	fromAddy := c.Query("fromaddy")
	toAddy := c.Query("toaddy")

	if fromAddy == "" || toAddy == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	rides, resp := app.Calls.SpecRides(fromAddy, toAddy)

	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": false, "error": resp.Err})
		return
	}

	c.JSON(200, gin.H{"msg": resp.Msg, "ok": true, "error": nil, "rides": rides})
	return
}

func (app *GuestRemote) AreaRides(c *gin.Context) {
	areaName := c.Query("areaName")

	if areaName == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	rides, resp := app.Calls.AreaRides(areaName)

	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": false, "error": resp.Err})
		return
	}

	c.JSON(200, gin.H{"msg": resp.Msg, "ok": true, "error": nil, "rides": rides})
	return
}

func (app *GuestRemote) CityRides(c *gin.Context) {
	cityName := c.Query("cityName")

	if cityName == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	rides, resp := app.Calls.CityRides(cityName)

	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": false, "error": resp.Err})
		return
	}

	c.JSON(200, gin.H{"msg": resp.Msg, "ok": true, "error": nil, "rides": rides})
	return
}

func (app *GuestRemote) StateRides(c *gin.Context) {
	stateName := c.Query("stateName")

	if stateName == "" {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	rides, resp := app.Calls.CityRides(stateName)

	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "ok": false, "error": resp.Err})
		return
	}

	c.JSON(200, gin.H{"msg": resp.Msg, "ok": true, "error": nil, "rides": rides})
	return
}
