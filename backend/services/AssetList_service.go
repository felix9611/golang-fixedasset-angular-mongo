package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
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

		assetItem.Status = 1
		assetItem.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		assetItem.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

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

func UpdateAssetItem(updateData *models.AssetLists) (interface{}, error) {
	collection := config.GetCollection("asset_lists")

	// Check if the asset item exists
	var existingAsset models.AssetLists
	err := collection.FindOne(context.Background(), bson.M{"_id": updateData.ID}).Decode(&existingAsset)
	if err != nil {
		return nil, err
	}

	if existingAsset.Status == 1 {

		updateData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

		res, err2 := collection.UpdateOne(context.Background(), bson.M{"_id": updateData.ID}, bson.M{"$set": updateData})
		if err2 != nil {
			return nil, err2
		}
		return res, nil
	} else {
		return "This asset item may be invalidated or not exist! Please contact admin!", nil
	}
}

func ListAssetItems(req *dto.ListAssetReqDto) (interface{}, error) {
	collection := config.GetCollection("asset_lists")
	//var assetItems []models.AssetLists

	filters := bson.M{}
	if req.AssetCode != "" {
		filters["asset_code"] = req.AssetCode
	}
	if req.AssetName != "" {
		filters["asset_name"] = req.AssetName
	}
	if len(req.TypeIds) > 0 {
		filters["type_ids"] = bson.M{"$in": req.TypeIds}
	}
	if len(req.PlaceIds) > 0 {
		filters["place_ids"] = bson.M{"$in": req.PlaceIds}
	}
	if len(req.DeptIds) > 0 {
		filters["dept_ids"] = bson.M{"$in": req.DeptIds}
	}
	if len(req.PurchaseDates) > 0 {
		filters["purchase_dates"] = bson.M{"$gte": req.PurchaseDates[0], "$lte": req.PurchaseDates[1]}
	}

	skip := int64((req.Page - 1) * req.Limit)
	limit := int64(req.Limit)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filters}}, // filters 必須係 bson.D or bson.M

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

		// $lookup assettypes
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "asset_types"},
				{Key: "let", Value: bson.D{
					{Key: "typeIdStr", Value: bson.D{
						{Key: "$convert", Value: bson.D{
							{Key: "input", Value: "$typeId"},
							{Key: "to", Value: "objectId"},
							{Key: "onError", Value: nil},
							{Key: "onNull", Value: nil},
						}},
					}},
				}},
				{Key: "pipeline", Value: mongo.Pipeline{
					{{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$eq", Value: bson.A{"$_id", "$$typeIdStr"}},
						}},
					}}},
				}},
				{Key: "as", Value: "assettype"},
			},
		}},


		// $addFields assetCodeInt
		{{
			Key: "$addFields", Value: bson.D{
				{Key: "assetCodeInt", Value: bson.D{
					{Key: "$convert", Value: bson.D{
						{Key: "input", Value: "$assetCode"},
						{Key: "to", Value: "int"},
						{Key: "onError", Value: 0}, // 防止 "" 出錯
						{Key: "onNull", Value: 0},
					}},
				}},
			},
		}},

		// unwind
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$location"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$department"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$assettype"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},

		// sort
		{{Key: "$sort", Value: bson.D{{Key: "assetCodeInt", Value: 1}}}},

		// pagination
		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: limit}},
	}



	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	count, errCount := collection.CountDocuments(ctx, filters)
	if errCount != nil {
		return nil, errCount
	}

	return gin.H{"lists": results, "total": count, "page": req.Page, "limit": req.Limit}, nil
}