package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"time"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    // "errors"
    //"log"
)

func CreateDepartment(department *models.Department) (interface{}, error) {
	filter := bson.M{"status": 1, "deptCode": department.DeptCode, "deptName": department.DeptName}

	// Check if a Department with the same name already exists
	collection := config.GetCollection("departments")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if department.Status == 0 {
			department.Status = 1
		}

		if department.CreatedAt.IsZero() {
			department.CreatedAt = time.Now()
		}

		if department.UpdatedAt.IsZero() {
			department.UpdatedAt = time.Now()
		}

		result, err := collection.InsertOne(ctx, department)
		if err != nil {
			return nil, err
		}
		return result.InsertedID, nil
	} else {
		return "Department with the same name already exists", nil
	}
}

func GetOneDepartment(id string) (*models.Department, error) {
	collection := config.GetCollection("departments")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var department models.Department

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&department)
	if err != nil {
		return nil, err
	}

	return &department, nil
}
