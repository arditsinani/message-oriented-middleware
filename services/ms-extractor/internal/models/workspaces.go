package models

import (
	"mom/services/ms-extractor/internal/db"
	"time"

)

type Workspace struct {
	ID          db.ObjectID		   `updatable:"false" unique:"true" db:"_id" json:"id" bson:"_id"`
	Name 		string             `updatable:"true" unique:"false" db:"name" json:"name" bson:"name"`
	DefaultLang string			   `updatable:"true" unique:"false" db:"default_lang" json:"default_lang" bson:"default_lang"`
	SSOS        []string 		   `updatable:"true" unique:"false" db:"ssos" json:"ssos" bson:"ssos"`
	Licenses    int                `updatable:"true" unique:"false" db:"licenses" json:"licenses" bson:"licenses"`
	Since   	time.Time          `updatable:"true" unique:"false" db:"since" json:"since" bson:"since"`
	Until   	time.Time          `updatable:"true" unique:"false" db:"created_at" json:"created_at" bson:"until"`
	Theme   	Theme          	   `updatable:"true" unique:"false" db:"colors" json:"colors" bson:"colors"`
	CreatedAt   time.Time          `updatable:"false" unique:"false" db:"created_at" json:"created_at" bson:"created_at" value:"NOW()"`
	UpdatedAt   time.Time          `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at" value:"NOW()"`
}

type Theme struct {
	Logo 	  	string 			   `updatable:"true" unique:"false" db:"name" json:"name" bson:"name"`
	Display	  	string 			   `updatable:"true" unique:"false" db:"display" json:"display" bson:"display"`
	Colors 	  	map[string]string  `updatable:"true" unique:"false" db:"colors" json:"colors" bson:"colors"`
}

type CreateWorkspaceForm struct {
	Name 		string             `updatable:"true" unique:"false" db:"name" json:"name" bson:"name"`
	DefaultLang string			   `updatable:"true" unique:"false" db:"default_lang" json:"default_lang" bson:"default_lang"`
	SSOS        []string 		   `updatable:"true" unique:"false" db:"ssos" json:"ssos" bson:"ssos"`
	Licenses    int                `updatable:"true" unique:"false" db:"licenses" json:"licenses" bson:"licenses"`
	Since   	time.Time          `updatable:"true" unique:"false" db:"since" json:"since" bson:"since"`
	Until   	time.Time          `updatable:"true" unique:"false" db:"created_at" json:"created_at" bson:"until"`
	Theme   	Theme          	   `updatable:"true" unique:"false" db:"colors" json:"colors" bson:"colors"`
	CreatedAt   time.Time          `updatable:"false" unique:"false" db:"created_at" json:"created_at" bson:"created_at" value:"NOW()"`
	UpdatedAt   time.Time          `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at" value:"NOW()"`
}

type UpdateWorkspaceForm struct {
	Name 		string             `updatable:"true" unique:"false" db:"name" json:"name" bson:"name"`
	DefaultLang string			   `updatable:"true" unique:"false" db:"default_lang" json:"default_lang" bson:"default_lang"`
	SSOS        []string 		   `updatable:"true" unique:"false" db:"ssos" json:"ssos" bson:"ssos"`
	Licenses    int                `updatable:"true" unique:"false" db:"licenses" json:"licenses" bson:"licenses"`
	Since   	time.Time          `updatable:"true" unique:"false" db:"since" json:"since" bson:"since"`
	Until   	time.Time          `updatable:"true" unique:"false" db:"created_at" json:"created_at" bson:"until"`
	Theme   	Theme          	   `updatable:"true" unique:"false" db:"colors" json:"colors" bson:"colors"`
	UpdatedAt   time.Time          `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at" value:"NOW()"`
}

type WorkspacesCollection struct {
	Collection []Workspace
}

const (
	WORKSPACESCOLLECTION = "workspaces"
)
