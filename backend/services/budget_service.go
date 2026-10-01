package services

import (
	"context"
	"fmt"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/models"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateBudgetRecord(budgetRecord *models.Budgets) (interface{}, error) {

	filter := bson.M{
		"deptId":     budgetRecord.DeptID,
		"placeId":    budgetRecord.PlaceId,
		"budgetNo":   budgetRecord.BudgetNo,
		"budgetName": budgetRecord.BudgetName,
		"year":       budgetRecord.Year,
		"month":      budgetRecord.Month,
	}

	collection := config.GetCollection("budgets")

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		if budgetRecord.Status == 0 {
			budgetRecord.Status = 1
		}

		if budgetRecord.CreatedAt.IsZero() {
			budgetRecord.CreatedAt = time.Now()
		}

		if budgetRecord.UpdatedAt.IsZero() {
			budgetRecord.UpdatedAt = time.Now()
		}

		if budgetRecord.BudgetNo == "" {
			budgetRecord.BudgetNo = getRandom10Digit()
		}

		result, err := collection.InsertOne(context.Background(), budgetRecord)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return "This Budget Record with the same details already exists", nil
	}
}

func GetOneBudgetRecord(id string) (interface{}, error) {
	collection := config.GetCollection("budgets")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err

	}
	filter := bson.M{"_id": objectID}

	var budgetRecord models.Budgets
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&budgetRecord)
	if err != nil {
		return nil, err
	}

	if budgetRecord.Status == 0 {
		return "This Budget Record is inactive", nil
	} else {
		return &budgetRecord, nil
	}
}

func VoidBudgetRecord(id string) (interface{}, error) {
	collection := config.GetCollection("budgets")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}

	var budgetRecord models.Budgets
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&budgetRecord)
	if err != nil {
		return nil, err
	}

	if budgetRecord.Status == 0 {
		return "This Budget Record is already inactive", nil
	} else {
		update := bson.M{
			"$set": bson.M{
				"status":    0,
				"updatedAt": time.Now(),
			},
		}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, err
		}
		return "This Budget Record has been voided just now", nil
	}
}

func UpdateBudgetRecord(updatedData *models.Budgets) (interface{}, error) {
	collection := config.GetCollection("budgets")

	filter := bson.M{"_id": updatedData.ID}

	var existingRecord models.Budgets
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&existingRecord)
	if err != nil {
		return nil, err
	}

	if existingRecord.Status == 0 {
		return "This Budget Record is inactive", nil
	} else {
		updatedData.UpdatedAt = time.Now()
		result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": updatedData})
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func ListBudgetRecordsFilter(pageDto *dto.ListBudgetRecordsDto) (interface{}, error) {
	collection := config.GetCollection("budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"status": 1,
	}

	if pageDto.Name == "" {
		filter["budgetName"] = bson.M{"$regex": pageDto.Name, "$options": "i"}
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", -1}})
	cursor, err := collection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var budgetRecords []models.Budgets
	for cursor.Next(ctx) {
		var budgetRecord models.Budgets
		if err := cursor.Decode(&budgetRecord); err != nil {
			return nil, err
		}
		budgetRecords = append(budgetRecords, budgetRecord)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return budgetRecords, nil
}

func ListBudgetRecords(pageDto *dto.ListBudgetRecordsDto) (interface{}, error) {
	if pageDto.Page < 1 {
		pageDto.Page = 1
	}

	if pageDto.Limit < 1 {
		pageDto.Limit = 10
	}

	skip := (pageDto.Page - 1) * pageDto.Limit
	limit := pageDto.Limit

	collection := config.GetCollection("budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"status": 1,
	}

	if pageDto.Name == "" {
		filter["budgetName"] = bson.M{"$regex": pageDto.Name, "$options": "i"}
	}

	count, errCount := collection.CountDocuments(ctx, filter)
	if errCount != nil {
		return nil, errCount
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		// $lookup locations
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "locations"},
				{Key: "let", Value: bson.D{
					{Key: "placeIdStr", Value: bson.D{
						{Key: "$convert", Value: bson.D{
							{Key: "input", Value: "$placeId"},
							{Key: "to", Value: "objectId"},
							{Key: "onError", Value: nil},
							{Key: "onNull", Value: nil},
						}},
					}},
				}},
				{Key: "pipeline", Value: mongo.Pipeline{
					{{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$eq", Value: bson.A{"$_id", "$$placeIdStr"}},
						}},
					}}},
				}},
				{Key: "as", Value: "location"},
			},
		}},

		// $lookup departments
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "departments"},
				{Key: "let", Value: bson.D{
					{Key: "deptIdStr", Value: bson.D{
						{Key: "$convert", Value: bson.D{
							{Key: "input", Value: "$deptId"},
							{Key: "to", Value: "objectId"},
							{Key: "onError", Value: nil},
							{Key: "onNull", Value: nil},
						}},
					}},
				}},
				{Key: "pipeline", Value: mongo.Pipeline{
					{{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$eq", Value: bson.A{"$_id", "$$deptIdStr"}},
						}},
					}}},
				}},
				{Key: "as", Value: "department"},
			},
		}},

		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$location"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$department"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},

		{{Key: "$sort", Value: bson.D{{"created_at", -1}}}},
		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: limit}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return gin.H{"lists": results, "total": count, "page": pageDto.Page, "limit": pageDto.Limit}, nil
}

func getRandom10Digit() string {
	rand.Seed(time.Now().UnixNano())
	num := 1000000000 + rand.Int63n(9000000000) // ensure 10 digits
	return fmt.Sprintf("%d", num)
}

func GetBudgetSummary() (interface{}, error) {
	collection := config.GetCollection("budgets")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "year", Value: bson.D{{Key: "$year", Value: "$budgetFrom"}}},
				{Key: "month", Value: bson.D{{Key: "$month", Value: "$budgetFrom"}}},
			}},
			{Key: "budgetAmount", Value: bson.D{{Key: "$sum", Value: "$budgetAmount"}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "budgetAmount", Value: 1},
			{Key: "yearMonth", Value: bson.D{
				{Key: "$concat", Value: bson.A{
					bson.D{{Key: "$toString", Value: "$_id.year"}},
					"-",
					bson.D{{Key: "$dateToString", Value: bson.D{
						{Key: "format", Value: "%B"},
						{Key: "date", Value: bson.D{
							{Key: "$dateFromParts", Value: bson.D{
								{Key: "year", Value: "$_id.year"},
								{Key: "month", Value: "$_id.month"},
								{Key: "day", Value: 1},
							}},
						}},
					}}},
				}},
			}},
			{Key: "year", Value: "$_id.year"},
			{Key: "month", Value: "$_id.month"},
		}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "year", Value: 1},
			{Key: "month", Value: 1},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []dto.BudgetSummary
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
