package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pattapong1/app/ent"
	"github.com/pattapong1/app/ent/activity"
	"github.com/pattapong1/app/ent/club"
	"github.com/pattapong1/app/ent/clubtypes"
	"github.com/pattapong1/app/ent/location"
	"github.com/pattapong1/app/ent/user"
)

// ClubController defines the struct for the Club controller
type ClubController struct {
	client *ent.Client
	router gin.IRouter
}
type Club struct {
	CLUBNAME       string
	Clubtypes     int
	Activity      int
	Location      int
	User           int
}

// CreateClub handles POST requests for adding Club entities
// @Summary Create Club
// @Description Create Club
// @ID create-Club
// @Accept   json
// @Produce  json
// @Param club body ent.Club true "Club entity"
// @Success 200 {object} ent.Club
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Clubs [post]
func (ctl *ClubController) CreateClub(c *gin.Context) {
	obj := Club{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Club binding failed",
		})
		return
	}
    l, err := ctl.client.Location.
		Query().
		Where(location.IDEQ(int(obj.Location))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Location not found",
		})
		return
	}
	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "User not found",
		})
		return
	}
	a, err := ctl.client.Activity.
		Query().
		Where(activity.IDEQ(int(obj.Activity))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Activity not found",
		})
		return
	}

	t, err := ctl.client.ClubTypes.
		Query().
		Where(clubtypes.IDEQ(int(obj.Clubtypes))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Types not found",
		})
		return
	}
	
	
	cl, err := ctl.client.Club.
		Create().
		SetCLUBNAME(obj.CLUBNAME).
		SetClubtypes(t).
		SetActivity(a).
		SetLocation(l).
		SetUser(u).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   cl,
	})
}

// GetClub handles GET requests to retrieve a Club entity
// @Summary Get a Club entity by ID
// @Description get Club by ID
// @ID get-Club
// @Produce  json
// @Param id path int true "Club ID"
// @Success 200 {object} ent.Club
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Clubs/{id} [get]
func (ctl *ClubController) GetClub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cl, err := ctl.client.Club.
		Query().
		Where(club.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cl)
}

// ListClub handles request to get a list of Club entities
// @Summary List Club entities
// @Description list Club entities
// @ID list-Club
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Club
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Clubs [get]
func (ctl *ClubController) ListClub(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	clubs, err := ctl.client.Club.
		Query().
		WithClubtypes().
		WithActivity().
		WithLocation().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clubs)
}

// DeleteClub handles DELETE requests to delete a Club entity
// @Summary Delete a Club entity by ID
// @Description get Club by ID
// @ID delete-Club
// @Produce  json
// @Param id path int true "Club ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Clubs/{id} [delete]
func (ctl *ClubController) DeleteClub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Club.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateClub handles PUT requests to update a Club entity
// @Summary Update a Club entity by ID
// @Description update Club by ID
// @ID update-Club
// @Accept   json
// @Produce  json
// @Param id path int true "Club ID"
// @Param Club body ent.Club true "Club entity"
// @Success 200 {object} ent.Club
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Clubs/{id} [put]
func (ctl *ClubController) UpdateClub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Club{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Club binding failed",
		})
		return
	}
	obj.ID = int(id)
	cl, err := ctl.client.Club.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cl)
}

// NewClubController creates and registers handles for the Club controller
func NewClubController(router gin.IRouter, client *ent.Client) *ClubController {
	clc := &ClubController{
		client: client,
		router: router,
	}
	clc.register()
	return clc
}

// InitClubController registers routes to the main engine
func (ctl *ClubController) register() {
	Clubs := ctl.router.Group("/Clubs")

	Clubs.GET("", ctl.ListClub)

	// CRUD
	Clubs.POST("", ctl.CreateClub)
	Clubs.GET(":id", ctl.GetClub)
	Clubs.PUT(":id", ctl.UpdateClub)
	Clubs.DELETE(":id", ctl.DeleteClub)
}
