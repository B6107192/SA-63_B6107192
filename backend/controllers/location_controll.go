package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pattapong1/app/ent"
	"github.com/pattapong1/app/ent/location"
)

// LocationController defines the struct for the Location controller
type LocationController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateLocation handles POST requests for adding Location entities
// @Summary Create Location
// @Description Create Location
// @ID create-Location
// @Accept   json
// @Produce  json
// @Param Location body ent.Location true "Location entity"
// @Success 200 {object} ent.Location
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Locations [post]
func (ctl *LocationController) CreateLocation(c *gin.Context) {
	obj := ent.Location{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Location binding failed",
		})
		return
	}

	l, err := ctl.client.Location.
		Create().
		SetCLUBELOCATIONNAME(obj.CLUBELOCATIONNAME).
		SetCLUBELOCATIONADDRESS(obj.CLUBELOCATIONADDRESS).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, l)
}

// GetLocation handles GET requests to retrieve a Location entity
// @Summary Get a Location entity by ID
// @Description get Location by ID
// @ID get-Location
// @Produce  json
// @Param id path int true "Location ID"
// @Success 200 {object} ent.Location
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Locations/{id} [get]
func (ctl *LocationController) GetLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	l, err := ctl.client.Location.
		Query().
		Where(location.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, l)
}

// ListLocation handles request to get a list of Location entities
// @Summary List Location entities
// @Description list Location entities
// @ID list-Location
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Location
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Locations [get]
func (ctl *LocationController) ListLocation(c *gin.Context) {
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

	Locations, err := ctl.client.Location.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Locations)
}

// DeleteLocation handles DELETE requests to delete a Location entity
// @Summary Delete a Location entity by ID
// @Description get Location by ID
// @ID delete-Location
// @Produce  json
// @Param id path int true "Location ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Locations/{id} [delete]
func (ctl *LocationController) DeleteLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Location.
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

// UpdateLocation handles PUT requests to update a Location entity
// @Summary Update a Location entity by ID
// @Description update Location by ID
// @ID update-Location
// @Accept   json
// @Produce  json
// @Param id path int true "Location ID"
// @Param Location body ent.Location true "Location entity"
// @Success 200 {object} ent.Location
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Locations/{id} [put]
func (ctl *LocationController) UpdateLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Location{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Location binding failed",
		})
		return
	}
	obj.ID = int(id)
	l, err := ctl.client.Location.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, l)
}

// NewLocationController creates and registers handles for the Location controller
func NewLocationController(router gin.IRouter, client *ent.Client) *LocationController {
	lc := &LocationController{
		client: client,
		router: router,
	}
	lc.register()
	return lc
}

// InitLocationController registers routes to the main engine
func (ctl *LocationController) register() {
	Locations := ctl.router.Group("/Locations")

	Locations.GET("", ctl.ListLocation)

	// CRUD
	Locations.POST("", ctl.CreateLocation)
	Locations.GET(":id", ctl.GetLocation)
	Locations.PUT(":id", ctl.UpdateLocation)
	Locations.DELETE(":id", ctl.DeleteLocation)
}
