package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"time"
)

func GetOneAssetItemByID(id string) (interface{}, error) {
	collection := config.GetCollection("asset_lists")
	var assetItem models.AssetLists

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}


	err2 := collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&assetItem)
	if err2 != nil {
		return nil, err2
	}
	if assetItem.Status == 0 {
		return "This Asset item is write off", nil
	} else {
		return &assetItem, nil
	}
}

func GetOneAssetItemByAssetCode(assetCode string) (interface{}, error) {
	collection := config.GetCollection("asset_lists")
	var assetItem models.AssetLists

	err := collection.FindOne(context.Background(), bson.M{"asset_code": assetCode}).Decode(&assetItem)
	if err != nil {
		return nil, err
	}
	if assetItem.Status == 0 {
		return "This Asset item is write off", nil
	} else {
		return &assetItem, nil
	}
}

func CreateAssetItem(assetItem *models.AssetLists) (interface{}, error) {
	collection := config.GetCollection("asset_lists")

	
	filter := bson.M{"assetName": assetItem.AssetName}

	var existingAsset models.AssetLists
	err := collection.FindOne(context.Background(), filter).Decode(&existingAsset)
	if err != nil {
		return nil, err
	}

	if existingAsset.AssetCode != "" && existingAsset.Status == 1 {
		return "This asset name already exist! Please check again!", nil
	} else {
		newAssetCode, err := CreateNewAssetCode()
		if err != nil {
			return nil, err
		}

		assetItem.AssetCode = newAssetCode
		_, errLast := collection.InsertOne(context.Background(), assetItem)
		return assetItem, errLast
	}
}

func formatNumber(num int, digits int) string {
	return fmt.Sprintf("%0*d", digits, num+1)
}

func CreateNewAssetCode() (string, error) {
	collection := config.GetCollection("asset_lists")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$addFields", Value: bson.D{
			{Key: "assetCodeInt", Value: bson.D{
				{Key: "$convert", Value: bson.D{
					{Key: "input", Value: "$assetCode"},
					{Key: "to", Value: "int"},
					{Key: "onError", Value: 0},
					{Key: "onNull", Value: 0},
				}},
			}},
		}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "maxNumber", Value: bson.D{{Key: "$max", Value: "$assetCodeInt"}}},
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

	// 格式化
	return formatNumber(maxNumber, 6), nil
}