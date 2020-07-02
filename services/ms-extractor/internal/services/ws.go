// workspaces service
package services

import (
	"log"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/models"

	"golang.org/x/net/context"
)

type WSS struct {
	DB *db.DB
}

func (s *WSS) Create(ctx context.Context, ws models.CreateWorkspaceForm, coll string) (models.Workspace, error) {
	result, err := s.DB.Create(ctx, ws, coll)
	if err != nil {
		return models.Workspace{}, err
	}
	response := models.Workspace{}
	id := result.InsertedID.(db.ObjectID)
	response.ID = id
	response.Name = ws.Name
	response.DefaultLang = ws.DefaultLang
	response.SSOS = ws.SSOS
	response.Licenses = ws.Licenses
	response.Since = ws.Since
	response.Until = ws.Until
	response.Theme = ws.Theme
	response.CreatedAt = ws.CreatedAt
	response.UpdatedAt = ws.UpdatedAt
	return response, nil
}

func (s *WSS) Get(ctx context.Context, filter db.MType, coll string) ([]*models.Workspace, error) {
	findOptions := s.DB.GetFindOptions()
	// TODO remove if not needed
	findOptions.SetLimit(1000)
	cur, err := s.DB.GetCursor(ctx, filter, coll, findOptions)
	var results []*models.Workspace
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Workspace
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

func (s *WSS) GetById(ctx context.Context, id db.ObjectID, coll string) (models.Workspace, error) {
	ws := models.Workspace{}
	singleResult := s.DB.GetById(ctx, id, coll)
	err := singleResult.Decode(&ws)
	return ws, err
}

func (s *WSS) Update(ctx context.Context, id db.ObjectID, ws models.UpdateWorkspaceForm, coll string) (*db.UpdateResult, error) {
	return s.DB.Update(ctx, id, ws, coll)
}

func (s *WSS) Delete(ctx context.Context, id db.ObjectID, coll string) (*db.DeleteResult, error) {
	return s.DB.Delete(ctx, id, coll)
}
