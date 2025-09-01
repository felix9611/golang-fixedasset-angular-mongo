package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/dto"
	"time"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
)

func  CreateCodeType(codeType *models.CodeTypes) (interface{}, error) {
	filter := bson.M{"status": 1, "valueCode": codeType.ValueCode, "valueName": codeType.ValueName}

	collection := config.GetCollection("code_types")

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {

		if codeType.Status == 0 {
			codeType.Status = 1
		}

		if codeType.CreatedAt.IsZero() {
			codeType.CreatedAt = time.Now()
		}

		if codeType.UpdatedAt.IsZero() {
			codeType.UpdatedAt = time.Now()
		}

		result, err := collection.InsertOne(context.Background(), codeType)
		if err != nil {
			return nil, err
		}

		return result, nil
	} else {
		return "CodeType with the same name already exists", nil
	}
}

func GetOneCodeType(id string) (interface{}, error) {
	collection := config.GetCollection("code_types")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID, "status": 1}

	var codeType models.CodeTypes

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&codeType)
	if err != nil {
		return nil, err
	}

	if codeType.Status == 0 {
		return "This Code Type is inactive", nil
	} else {
		return &codeType, nil
	}
}

func VoidOneCodeType(id string) (interface{}, error) {
	collection := config.GetCollection("code_types")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	
	filter := bson.M{"_id": objectID}

	var codeType models.CodeTypes

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&codeType)

	if err != nil {
		return nil, err
	}

	if codeType.Status == 1 {

		result, err := collection.UpdateOne(ctx, filter, bson.M{
            "$set": bson.M{
                "status": 0,
                "updated_at": time.Now(),
            },
        })

        if err != nil {
            return nil, err
        }
        return result, nil
	} else {
		return "This Code Type is already voided", nil
	}
}

func UpdateCodeType(updateData *models.CodeTypes) (interface{}, error) {
	collection := config.GetCollection("code_types")

	filter := bson.M{"_id": updateData.ID, "status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var checkCodeType models.CodeTypes
	err := collection.FindOne(ctx, filter).Decode(&checkCodeType)
	if err != nil {
		return nil, err
	}

	if checkCodeType.Status == 1 {
		updateData.UpdatedAt = time.Now()

		result, err := collection.UpdateOne(ctx, filter, bson.M{
			"$set": updateData,
		})
		if err != nil {
			return nil, err
		}

		return result, nil
	} else {
		return "This Code Type is already voided", nil
	}
}

func CodeTypeList(pageDto *dto.CodeTypeListDto) (interface{}, error) {

	if pageDto.Page < 1 {
		pageDto.Page = 1
	}

	if pageDto.Limit < 1 {
		pageDto.Limit = 10
	}


	collection := config.GetCollection("code_types")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

	filter := bson.M{"status": 1}

	if pageDto.Name != "" {
		filter = bson.M{
			"status": 1,
			"$or": []bson.M{
				{"type": bson.M{"$regex": pageDto.Name, "$options": "i"}},
				{"valueCode": bson.M{"$regex": pageDto.Name, "$options": "i"}},
				{"valueName": bson.M{"$regex": pageDto.Name, "$options": "i"}},
			},
		}
	}

	count, errCount := collection.CountDocuments(ctx, filter)
	if errCount != nil {
		return nil, errCount
	}

	skip := int64((pageDto.Page - 1) * pageDto.Limit)
    limit := int64(pageDto.Limit)

	findOptions := options.Find()
    findOptions.SetSkip(skip)
    findOptions.SetLimit(limit)
    findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var codeTypes []models.CodeTypes
	for cursor.Next(ctx) {
		var codeType models.CodeTypes
		if err := cursor.Decode(&codeType); err != nil {
			return nil, err
		}
		codeTypes = append(codeTypes, codeType)
	}

	if err := cursor.Err(); err != nil {
        return nil, err
    }

	return gin.H{"lists": codeTypes, "total": count, "page": pageDto.Page, "limit": pageDto.Limit }, nil
}

func ListCodeTypeByType(typeString string) (interface{}, error) {
	collection := config.GetCollection("code_types")

	filter := bson.M{"type": typeString, "status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var codeTypes []models.CodeTypes
	for cursor.Next(ctx) {
		var codeType models.CodeTypes
		if err := cursor.Decode(&codeType); err != nil {
			return nil, err
		}
		codeTypes = append(codeTypes, codeType)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return gin.H{"datas": codeTypes}, nil
}