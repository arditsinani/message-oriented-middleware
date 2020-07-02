// test controller
package controllers

import (
	"context"
	"fmt"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/models"
	"mom/services/ms-extractor/internal/services"

	"github.com/gin-gonic/gin"
)

type TestCtrl struct {
	Config  *config.Config
	DB      *db.DB
	Service services.TestS
}

func (ctrl *TestCtrl) Create(c *gin.Context) {
	// validate payload
	var createTestForm models.CreateTestForm
	if err := c.ShouldBindJSON(&createTestForm); err != nil {
		c.Error(err)
		return
	}
	inserted, err := ctrl.Service.Create(context.TODO(), createTestForm, models.TESTCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, inserted)
}

func (ctrl *TestCtrl) Get(c *gin.Context) {
	filter := db.MType{}
	tests, err := ctrl.Service.Get(context.TODO(), filter, models.TESTCOLLECTION)
	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}
	c.JSON(200, tests)
}

func (ctrl *TestCtrl) GetById(c *gin.Context) {
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
	res, err := ctrl.Service.GetById(context.TODO(), id, models.TESTCOLLECTION)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}

func (ctrl *TestCtrl) Update(c *gin.Context) {
	id, err := ctrl.DB.GetObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	var updateTestForm models.UpdateTestForm
	if err := c.ShouldBindJSON(&updateTestForm); err != nil {
		c.Error(err)
		return
	}
	updated, err := ctrl.Service.Update(context.TODO(), id, updateTestForm, models.TESTCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, updated)
}
func (ctrl *TestCtrl) Delete(c *gin.Context) {
	id, err := ctrl.DB.GetObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	test, err := ctrl.Service.Delete(context.TODO(), id, models.TESTCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, test)
}
