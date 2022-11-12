package controllers

import (
	"be-shareit/models"
	"be-shareit/services"

	"github.com/gin-gonic/gin"
)

type AdminRemote struct {
	Calls services.AdminContracts
}

func ARemoteMaker(s services.AdminContracts) *AdminRemote {
	return &AdminRemote{
		Calls: s,
	}
}

func (app *AdminRemote) AdminRoutes(incomingRoutes *gin.RouterGroup) {
	routes := incomingRoutes.Group("admin")
	routes.GET("/allrides", app.ShowRides)
	routes.GET("/allusers", app.ShowUsers)
	routes.PUT("/edituser", app.EditUser)
	routes.DELETE("/rmusers", app.DeleteUsers)
	routes.DELETE("/rmrides", app.DeleteRides)
}
func (app *AdminRemote) ShowRides(c *gin.Context) {
	rides, resp := app.Calls.ShowRides()
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": resp.Err.Error()})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": rides, "error": nil})
	return
}

func (app *AdminRemote) ShowUsers(c *gin.Context) {
	users, resp := app.Calls.ShowUsers()
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": resp.Err.Error()})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": users, "error": nil})
	return
}

func (app *AdminRemote) EditUser(c *gin.Context) {

	var changedUser models.User
	if err := c.BindJSON(&changedUser); err != nil {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.EditUser(changedUser)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": resp.Err.Error()})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": nil})
	return
}

func (app *AdminRemote) DeleteUsers(c *gin.Context) {

	var userids []string
	if err := c.BindJSON(&userids); err != nil {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.DeleteUsers(userids)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": resp.Err.Error()})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": nil})
	return

}

func (app *AdminRemote) DeleteRides(c *gin.Context) {
	var rideids []string
	if err := c.BindJSON(&rideids); err != nil {
		c.JSON(403, gin.H{"msg": "Parse Error", "ok": false, "error": "Bad request"})
		return
	}

	resp := app.Calls.DeleteRides(rideids)
	if resp.Err != nil {
		c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": resp.Err.Error()})
		return
	}

	c.JSON(resp.Stat, gin.H{"msg": resp.Msg, "data": nil, "error": nil})
	return

}
