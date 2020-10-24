package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pattapong1/app/ent"
	"github.com/pattapong1/app/ent/activity"
)

// ActivityController defines the struct for the Activity controller
type ActivityController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateActivity handles POST requests for adding Activity entities
// @Summary Create Activity
// @Description Create Activity
// @ID create-Activity
// @Accept   json
// @Produce  json
// @Param Activity body ent.Activity true "Activity entity"
// @Success 200 {object} ent.Activity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitys [post]
func (ctl *ActivityController) CreateActivity(c *gin.Context) {
	obj := ent.Activity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Activity binding failed",
		})
		return
	}

	a, err := ctl.client.Activity.
		Create().
		SetCLUBEACTIVITYNAME(obj.CLUBEACTIVITYNAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, a)
}

// GetActivity handles GET requests to retrieve a Activity entity
// @Summary Get a Activity entity by ID
// @Description get Activity by ID
// @ID get-Activity
// @Produce  json
// @Param id path int true "Activity ID"
// @Success 200 {object} ent.Activity
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitys/{id} [get]
func (ctl *ActivityController) GetActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	a, err := ctl.client.Activity.
		Query().
		Where(activity.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, a)
}

// ListActivity handles request to get a list of Activity entities
// @Summary List Activity entities
// @Description list Activity entities
// @ID list-Activity
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Activity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitys [get]
func (ctl *ActivityController) ListActivity(c *gin.Context) {
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

	Activitys, err := ctl.client.Activity.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Activitys)
}

// DeleteActivity handles DELETE requests to delete a Activity entity
// @Summary Delete a Activity entity by ID
// @Description get Activity by ID
// @ID delete-Activity
// @Produce  json
// @Param id path int true "Activity ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitys/{id} [delete]
func (ctl *ActivityController) DeleteActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Activity.
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

// UpdateActivity handles PUT requests to update a Activity entity
// @Summary Update a Activity entity by ID
// @Description update Activity by ID
// @ID update-Activity
// @Accept   json
// @Produce  json
// @Param id path int true "Activity ID"
// @Param Activity body ent.Activity true "Activity entity"
// @Success 200 {object} ent.Activity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitys/{id} [put]
func (ctl *ActivityController) UpdateActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Activity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Activity binding failed",
		})
		return
	}
	obj.ID = int(id)
	a, err := ctl.client.Activity.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, a)
}

// NewActivityController creates and registers handles for the Activity controller
func NewActivityController(router gin.IRouter, client *ent.Client) *ActivityController {
	ac := &ActivityController{
		client: client,
		router: router,
	}
	ac.register()
	return ac
}

// InitActivityController registers routes to the main engine
func (ctl *ActivityController) register() {
	Activitys := ctl.router.Group("/Activitys")

	Activitys.GET("", ctl.ListActivity)

	// CRUD
	Activitys.POST("", ctl.CreateActivity)
	Activitys.GET(":id", ctl.GetActivity)
	Activitys.PUT(":id", ctl.UpdateActivity)
	Activitys.DELETE(":id", ctl.DeleteActivity)
}
