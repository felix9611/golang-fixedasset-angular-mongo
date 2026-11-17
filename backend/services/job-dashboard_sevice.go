package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetJobDashboardCollections() *mongo.Collection {
	return config.GetCollection("job_dashboards")
}

func CreateTasks(jobDashboard *models.JobDashboard) (interface{}, error) {
	collection := GetJobDashboardCollections()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"title": jobDashboard.Title,
	}

	var existingJobDashboard models.JobDashboard
	err := collection.FindOne(ctx, filter).Decode(&existingJobDashboard)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			newNumber, err := CreateJobNumber()
			if err != nil {
				return nil, err
			}

			jobDashboard.Number = newNumber
			jobDashboard.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			jobDashboard.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			_, errCr := collection.InsertOne(ctx, jobDashboard)
			if errCr != nil {
				return nil, errCr
			}

			CreateActionRecord("Job Dashboard Create", "POST", "Job Dashboard", jobDashboard, "Success")
			return gin.H{"message": "Job Dashboard created successfully"}, nil
		} else {
			CreateActionRecord("Job Dashboard Create", "POST", "Job Dashboard", jobDashboard, "Failed")
			return gin.H{"error": "Job Dashboard already exists"}, nil
		}
	}

	CreateActionRecord("Job Dashboard Create", "POST", "Job Dashboard", jobDashboard, "Failed")
	return gin.H{"error": "Job Dashboard already exists"}, nil
}

func CreateJobNumber() (string, error) {
	collection := GetJobDashboardCollections()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$addFields", Value: bson.D{
			{Key: "numberInt", Value: bson.D{
				{Key: "$convert", Value: bson.D{
					{Key: "input", Value: "$number"},
					{Key: "to", Value: "int"},
					{Key: "onError", Value: 0},
					{Key: "onNull", Value: 0},
				}},
			}},
		}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "maxNumber", Value: bson.D{{Key: "$max", Value: "$numberInt"}}},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return "", err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return "", err
	}

	var maxNumber int
	if len(results) > 0 {
		if val, ok := results[0]["maxNumber"]; ok {
			// BSON to Go int
			switch v := val.(type) {
			case int32:
				maxNumber = int(v)
			case int64:
				maxNumber = int(v)
			case float64:
				maxNumber = int(v)
			default:
				maxNumber = 0
			}
		}
	}

	// format
	return formatNumber(maxNumber, 8), nil
}
