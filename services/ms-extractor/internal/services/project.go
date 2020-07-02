// project service
package services

import (
	"context"
	"log"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/models"
)

type PrjS struct {
	DB *db.DB
}

func (s *PrjS) Create(ctx context.Context, prj models.CreateProjectForm, coll string) (models.Project, error) {
	result, err := s.DB.Create(ctx, prj, coll)
	if err != nil {
		return models.Project{}, err
	}
	response := models.Project{}
	id := result.InsertedID.(db.ObjectID)
	response.ID = id
	response.Title = prj.Title
	response.Deleted = prj.Deleted
	response.WorkspaceId = prj.WorkspaceId
	response.CreatedAt = prj.CreatedAt
	response.UpdatedAt = prj.UpdatedAt
	return response, nil
}

func (s *PrjS) Get(ctx context.Context, filter db.MType, coll string) ([]*models.Project, error) {
	findOptions := s.DB.GetFindOptions()
	// TODO remove if not needed
	findOptions.SetLimit(1000)
	cur, err := s.DB.GetCursor(ctx, filter, coll, findOptions)
	var results []*models.Project
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Project
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err
}

func (s *PrjS) GetById(ctx context.Context, id db.ObjectID, coll string) (models.Project, error) {
	prj := models.Project{}
	singleResult := s.DB.GetById(ctx, id, coll)
	err := singleResult.Decode(&prj)
	return prj, err
}

func (s *PrjS) Update(ctx context.Context, id db.ObjectID, prj models.UpdateProjectForm, coll string) (*db.UpdateResult, error) {
	return s.DB.Update(ctx, id, prj, coll)
}

func (s *PrjS) Delete(ctx context.Context, id db.ObjectID, coll string) (*db.DeleteResult, error) {
	return s.DB.Delete(ctx, id, coll)
}
