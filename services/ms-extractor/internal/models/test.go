package models

import "mom/services/ms-extractor/internal/db"

type Test struct {
	ID 		db.ObjectID		`updatable:"false" unique:"true" db:"_id" json:"id" bson:"_id"`
	Name 	string			`updatable:"true" unique:"false" db:"name" json:"name" bson:"name"`
	Surname string			`updatable:"true" unique:"false" db:"surname" json:"surname" bson:"surname"`
}
type CreateTestForm struct {
	Name    string 			`json:"name" binding:"required"`
	Surname string 			`json:"surname" binding:"required"`
}

type UpdateTestForm struct {
	Name    	string `json:"name" binding:"required"`
	Surname  	string `json:"surname" binding:"required"`
}

type TestCollection struct {
	Collection []Test
}

const (
	TESTCOLLECTION = "test_collection"
)