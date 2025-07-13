package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"time"

//	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTongs(tongs *models.Tongs) (interface{}, error) {
	collection := config.GetCollection("tongs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, tongs)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
