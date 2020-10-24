package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pattapong1/app/ent"
	"github.com/pattapong1/app/ent/clubtypes"
)

// ClubTypesController defines the struct for the ClubTypes controller
type ClubTypesController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateClubTypes handles POST requests for adding ClubTypes entities
// @Summary Create ClubTypes
// @Description Create ClubTypes
// @ID create-ClubTypes
// @Accept   json
// @Produce  json
// @Param ClubTypes body ent.ClubTypes true "ClubTypes entity"
// @Success 200 {object} ent.ClubTypes
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ClubTypess [post]
func (ctl *ClubTypesController) CreateClubTypes(c *gin.Context) {
	obj := ent.ClubTypes{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "ClubTypes binding failed",
		})
		return
	}

	t, err := ctl.client.ClubTypes.
		Create().
		SetCLUBETYPENAME(obj.CLUBETYPENAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetClubTypes handles GET requests to retrieve a ClubTypes entity
// @Summary Get a ClubTypes entity by ID
// @Description get ClubTypes by ID
// @ID get-ClubTypes
// @Produce  json
// @Param id path int true "ClubTypes ID"
// @Success 200 {object} ent.ClubTypes
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ClubTypess/{id} [get]
func (ctl *ClubTypesController) GetClubTypes(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := ctl.client.ClubTypes.
		Query().
		Where(clubtypes.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListClubTypes handles request to get a list of ClubTypes entities
// @Summary List ClubTypes entities
// @Description list ClubTypes entities
// @ID list-ClubTypes
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ClubTypes
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ClubTypess [get]
func (ctl *ClubTypesController) ListClubTypes(c *gin.Context) {
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

	ClubTypes, err := ctl.client.ClubTypes.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ClubTypes)
}

// DeleteClubTypes handles DELETE requests to delete a ClubTypes entity
// @Summary Delete a ClubTypes entity by ID
// @Description get ClubTypes by ID
// @ID delete-ClubTypes
// @Produce  json
// @Param id path int true "ClubTypes ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ClubTypess/{id} [delete]
func (ctl *ClubTypesController) DeleteClubTypes(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ClubTypes.
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

// UpdateClubTypes handles PUT requests to update a ClubTypes entity
// @Summary Update a ClubTypes entity by ID
// @Description update ClubTypes by ID
// @ID update-ClubTypes
// @Accept   json
// @Produce  json
// @Param id path int true "ClubTypes ID"
// @Param ClubTypes body ent.ClubTypes true "ClubTypes entity"
// @Success 200 {object} ent.ClubTypes
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ClubTypess/{id} [put]
func (ctl *ClubTypesController) UpdateClubTypes(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ClubTypes{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "ClubTypes binding failed",
		})
		return
	}
	obj.ID = int(id)
	t, err := ctl.client.ClubTypes.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, t)
}

// NewClubTypesController creates and registers handles for the ClubTypes controller
func NewClubTypesController(router gin.IRouter, client *ent.Client) *ClubTypesController {
	tc := &ClubTypesController{
		client: client,
		router: router,
	}
	tc.register()
	return tc
}

// InitClubTypesController registers routes to the main engine
func (ctl *ClubTypesController) register() {
	ClubTypes := ctl.router.Group("/ClubTypess")

	ClubTypes.GET("", ctl.ListClubTypes)

	// CRUD
	ClubTypes.POST("", ctl.CreateClubTypes)
	ClubTypes.GET(":id", ctl.GetClubTypes)
	ClubTypes.PUT(":id", ctl.UpdateClubTypes)
	ClubTypes.DELETE(":id", ctl.DeleteClubTypes)
}
