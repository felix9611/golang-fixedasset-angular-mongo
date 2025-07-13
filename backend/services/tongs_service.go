package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"time"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func Foo() error {
    return errors.New("some error happened")
}

func CreateTongs(tongs *models.Tongs) (interface{}, error) {

    filter := bson.M{"status": 1, "name": tongs.Name}

    // Check if a Tong with the same name already exists
    collection := config.GetCollection("tongs")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    count, err := collection.CountDocuments(ctx, filter)
    if err != nil {
        return nil, err
    }

    if count == 0 {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        if tongs.Status == 0 {
            tongs.Status = 1
        }

        if tongs.CreatedAt.IsZero() {
            tongs.CreatedAt = time.Now()
        }

        if tongs.UpdatedAt.IsZero() {
            tongs.UpdatedAt = time.Now()
        }

        result, err := collection.InsertOne(ctx, tongs)
        if err != nil {
            return nil, err
        }
        return result.InsertedID, nil
    } else {
        return "Tong with the same name already exists", nil
    }

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

    if updateData.UpdatedAt.IsZero() {
        updateData.UpdatedAt = time.Now()
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
