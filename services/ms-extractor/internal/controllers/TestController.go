package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/models"
	"mom/services/ms-extractor/internal/services"
)

type TestController struct {
	Config config.Config
	Mongo *mongo.Client
}

func (ctrl *TestController) Create (c *gin.Context) {
	// validate payload
	var createTestForm models.CreateTestForm
	if err := c.ShouldBindJSON(&createTestForm); err != nil {
		c.Error(err)
		return
	}
	inserted, err := services.Insert(ctrl.Mongo,context.TODO(), createTestForm, ctrl.Config.Mongo.DatabaseName, models.TESTCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, inserted)
}

func (ctrl *TestController) Get (c *gin.Context) {
	filter := bson.M{}
	tests, err := services.Get(ctrl.Mongo, context.TODO(), filter, ctrl.Config.Mongo.DatabaseName, models.TESTCOLLECTION)
	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}
	c.JSON(200, tests)
}

func (ctrl *TestController) GetById (c *gin.Context) {
	if c.Param("id") == "" || len(c.Param("id")) != 24 {
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	test, err := services.GetById(ctrl.Mongo, context.TODO(), id, ctrl.Config.Mongo.DatabaseName, models.TESTCOLLECTION)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, test)
}

func (ctrl *TestController) Update (c *gin.Context) {
	// TODO validate hex id
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	var updateTestForm models.UpdateTestForm
	if err := c.ShouldBindJSON(&updateTestForm); err != nil {
		c.Error(err)
		return
	}
	updated, err := services.Update(ctrl.Mongo, context.TODO(), id, updateTestForm, ctrl.Config.Mongo.DatabaseName, models.TESTCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, updated)
}
func (ctrl *TestController) Delete (c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	test, err := services.Delete(ctrl.Mongo, context.TODO(), id, ctrl.Config.Mongo.DatabaseName, models.TESTCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, test)
}