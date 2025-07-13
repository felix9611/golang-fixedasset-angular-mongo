package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"time"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTongs(tongs *models.Tongs) (interface{}, error) {
	collection := config.GetCollection("tongs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    if tongs.Status == 0 {
		tongs.Status = 1
	}

	result, err := collection.InsertOne(ctx, tongs)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func GetOneTongsRun(id string) (*models.Tongs, error) {
	collection := config.GetCollection("tongs")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var tongs models.Tongs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&tongs)
	if err != nil {
		return nil, err
	}

	return &tongs, nil
}

func UpdateTongsByID(id string, updateData *models.Tongs) (interface{}, error) {
	collection := config.GetCollection("tongs")

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    filter := bson.M{"_id": objectID}

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    var checkTongs models.Tongs

    err = collection.FindOne(ctx, filter).Decode(&checkTongs)

    if err != nil {
        return nil, err
    }

    result, err := collection.UpdateOne(ctx, filter, bson.M{
        "$set": updateData,
    })

    if err != nil {
        return nil, err
    }

    if result.MatchedCount == 0 {
        return nil, mongo.ErrNoDocuments
    }
    return result.ModifiedCount, nil
}
