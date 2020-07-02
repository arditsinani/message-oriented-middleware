// project controller
package controllers

import (
	"context"
	"fmt"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/models"
	"mom/services/ms-extractor/internal/services"
	"time"

	"github.com/gin-gonic/gin"
)

type PrjCtrl struct {
	Config  *config.Config
	DB      *db.DB
	Service services.PrjS
}

func (ctrl *PrjCtrl) Create(c *gin.Context) {
	// validate payload
	var createPrjF models.CreateProjectForm
	if err := c.ShouldBindJSON(&createPrjF); err != nil {
		c.Error(err)
		c.JSON(500, "error parsing request body")
		return
	}
	createPrjF.CreatedAt = time.Now()
	createPrjF.UpdatedAt = time.Now()
	inserted, err := ctrl.Service.Create(context.TODO(), createPrjF, models.PROJECTSCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, inserted)
}

func (ctrl *PrjCtrl) Get(c *gin.Context) {
	filter := db.MType{}
	projects, err := ctrl.Service.Get(context.TODO(), filter, models.PROJECTSCOLLECTION)
	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}
	c.JSON(200, projects)
}

func (ctrl *PrjCtrl) GetById(c *gin.Context) {
	if c.Param("id") == "" || len(c.Param("id")) != 24 {
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}
	id, err := ctrl.DB.GetObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := ctrl.Service.GetById(context.TODO(), id, models.PROJECTSCOLLECTION)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}

func (ctrl *PrjCtrl) Update(c *gin.Context) {
	id, err := ctrl.DB.GetObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	var updatePrjF models.UpdateProjectForm
	if err := c.ShouldBindJSON(&updatePrjF); err != nil {
		c.Error(err)
		return
	}
	updatePrjF.UpdatedAt = time.Now()
	updated, err := ctrl.Service.Update(context.TODO(), id, updatePrjF, models.PROJECTSCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, updated)
}

func (ctrl *PrjCtrl) Delete(c *gin.Context) {
	id, err := ctrl.DB.GetObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	prj, err := ctrl.Service.Delete(context.TODO(), id, models.PROJECTSCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, prj)
}
