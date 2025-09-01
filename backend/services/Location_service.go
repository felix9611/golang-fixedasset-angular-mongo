package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/dto"
	"time"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
)

func CreateLocation(location *models.Locations) (interface{}, error) {
	collection := config.GetCollection("locations")

	filter := bson.M{"status": 1, "vendorName": location.PlaceCode, "placeCode": location.PlaceName  }

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return nil, err
	}

	if count == 0 { 
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if location.Status == 0 {
			location.Status = 1
		}

		if location.CreatedAt.IsZero() {
			location.CreatedAt = time.Now()
		}

		if location.UpdatedAt.IsZero() {
			location.UpdatedAt = time.Now()
		}

		_, err := collection.InsertOne(ctx, location)

		if err != nil {
			return nil, err
		}

		return location, nil

	} else {
		return "Location with the same name already exists", nil
	}
}

func GetOneLocationById(id string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("locations")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID, "status": 1}

	var location models.Locations
	err = collection.FindOne(ctx, filter).Decode(&location)

	if err != nil {
		return nil, err
	}

	if location.Status == 0 {
		return "Location is inactive", nil
	} else {
		return &location, nil
	}
}

func UpdateLocationById(updateData *models.Locations) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("locations")
	filter := bson.M{"_id": updateData.ID, "status": 1}

	var existingLocation models.Locations
	err := collection.FindOne(ctx, filter).Decode(&existingLocation)
	if err != nil {
		return nil, err
	}

	if existingLocation.Status == 1 {
		updateData.UpdatedAt = time.Now()

		result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": updateData})
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return "Location is inactive", nil
	}
}

func InactiveLocationByID(id string) (interface{}, error) {
	collection := config.GetCollection("locations")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID, "status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	var existingLocation models.Locations
	err = collection.FindOne(ctx, filter).Decode(&existingLocation)
	if err != nil {
		return nil, err
	}

	if existingLocation.Status == 1 {
		_, err := collection.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{
				"status":    0,
				"updatedAt": time.Now(),
			},
		})

		if err != nil {
			return nil, err
		}
		return "Location deactivated successfully", nil
	} else {
		return "Location is already inactive", nil
	}
}

func LocationList(pageDto *dto.LocationPageDto) (interface{}, error) {
	if pageDto.Page < 1 {
		pageDto.Page = 1
	}
	if pageDto.Limit < 1 {
		pageDto.Limit = 10
	}

	skip := (pageDto.Page - 1) * pageDto.Limit
	limit := pageDto.Limit

	collection := config.GetCollection("locations")
	filter := bson.M{"status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, errCount := collection.CountDocuments(ctx, filter)

	if errCount != nil {
		return nil, errCount
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{"created_at", -1}})
	cursor, err := collection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	var locations []models.Locations
	for cursor.Next(ctx) {
		var location models.Locations
		if err := cursor.Decode(&location); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return gin.H{"lists": locations, "total": count, "page": pageDto.Page, "limit": pageDto.Limit }, nil
}

func ListAllLocation() (interface{}, error) {
	collection := config.GetCollection("locations")

	filter := bson.M{"status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var locations []models.Locations
	for cursor.Next(ctx) {
		var location models.Locations
		if err := cursor.Decode(&location); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return gin.H{"data": locations}, nil
}
