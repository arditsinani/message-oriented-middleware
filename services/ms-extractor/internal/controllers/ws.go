// workspaces controller
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

type WSCtrl struct {
	Config  *config.Config
	DB      *db.DB
	Service services.WSS
}

func (ctrl *WSCtrl) Create(c *gin.Context) {
	// validate payload
	var createWsForm models.CreateWorkspaceForm
	if err := c.ShouldBindJSON(&createWsForm); err != nil {
		c.Error(err)
		c.JSON(500, "error parsing request body")
		return
	}
	createWsForm.CreatedAt = time.Now()
	createWsForm.UpdatedAt = time.Now()
	inserted, err := ctrl.Service.Create(context.TODO(), createWsForm, models.WORKSPACESCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, inserted)
}

func (ctrl *WSCtrl) Get(c *gin.Context) {
	filter := db.MType{}
	workspaces, err := ctrl.Service.Get(context.TODO(), filter, models.WORKSPACESCOLLECTION)
	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}
	c.JSON(200, workspaces)
}

func (ctrl *WSCtrl) GetById(c *gin.Context) {
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
	res, err := ctrl.Service.GetById(context.TODO(), id, models.WORKSPACESCOLLECTION)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}

func (ctrl *WSCtrl) Update(c *gin.Context) {
	id, err := ctrl.DB.GetObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	var updateWsForm models.UpdateWorkspaceForm
	if err := c.ShouldBindJSON(&updateWsForm); err != nil {
		c.Error(err)
		return
	}
	updateWsForm.UpdatedAt = time.Now()
	updated, err := ctrl.Service.Update(context.TODO(), id, updateWsForm, models.WORKSPACESCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, updated)
}

func (ctrl *WSCtrl) Delete(c *gin.Context) {
	id, err := ctrl.DB.GetObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	ws, err := ctrl.Service.Delete(context.TODO(), id, models.WORKSPACESCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, ws)
}
