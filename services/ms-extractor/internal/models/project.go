package models

import (
	"mom/services/ms-extractor/internal/db"
	"time"
)

type Project struct {
	ID          db.ObjectID `updatable:"false" unique:"true" db:"_id" json:"id" bson:"_id"`
	Title       string      `updatable:"true" unique:"false" db:"title" json:"title" bson:"title"`
	WorkspaceId db.ObjectID `updatable:"false" unique:"false" db:"workspace_id" json:"workspace_id" bson:"workspace_id"`
	Deleted     bool        `updatable:"true" unique:"false" db:"deleted" json:"deleted" bson:"deleted"`
	CreatedAt   time.Time   `updatable:"false" unique:"false" db:"created_at" json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time   `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at"`
}

type CreateProjectForm struct {
	Title       string      `updatable:"true" unique:"false" db:"title" json:"title" bson:"title"`
	WorkspaceId db.ObjectID `updatable:"false" unique:"false" db:"workspace_id" json:"workspace_id" bson:"workspace_id"`
	Deleted     bool        `updatable:"true" unique:"false" db:"deleted" json:"deleted" bson:"deleted"`
	CreatedAt   time.Time   `updatable:"false" unique:"false" db:"created_at" json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time   `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at"`
}

type UpdateProjectForm struct {
	Title       string      `updatable:"true" unique:"false" db:"title" json:"title" bson:"title"`
	WorkspaceId db.ObjectID `updatable:"false" unique:"false" db:"workspace_id" json:"workspace_id" bson:"workspace_id"`
	OwnerId     db.ObjectID `updatable:"false" unique:"false" db:"owner_id" json:"owner_id" bson:"owner_id"`
	Deleted     bool        `updatable:"true" unique:"false" db:"deleted" json:"deleted" bson:"deleted"`
	UpdatedAt   time.Time   `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at"`
}

type ProjectsCollection struct {
	Collection []Project
}

const (
	PROJECTSCOLLECTION = "projects"
)
