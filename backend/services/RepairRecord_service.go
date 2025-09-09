package services

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/dto"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRepairRecord(record *models.RepairRecords) (interface{}, error) {
	collection := config.GetCollection("repair_records")
	// assetCollection := config.GetCollection("asset_lists")

//	filters := bson.M{"_id": record.AssetId}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if record.Status == 0 {
			record.Status = 1
		}

		if record.CreatedAt.IsZero() {
			record.CreatedAt = time.Now()
		}

		if record.UpdatedAt.IsZero() {
			record.UpdatedAt = time.Now()
		}

		result, err := collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	CreateActionRecord("Repair Record Create", "POST", "Repair Record", record, "Success")

	return result, nil

/*	var asset *models.AssetLists

	err := assetCollection.FindOne(ctx, filters).Decode(&asset)
	if err != nil {
		return nil, err
	}

	if asset.Status == 0 {
		CreateActionRecord("Repair Record Create", "POST", "Repair Record", record, "Failed")
		return "This asset item is already written off", nil
	} else {
		if record.Status == 0 {
			record.Status = 1
		}

		if record.CreatedAt.IsZero() {
			record.CreatedAt = time.Now()
		}

		if record.UpdatedAt.IsZero() {
			record.UpdatedAt = time.Now()
		}

		result, err := collection.InsertOne(ctx, record)
		if err != nil {
			return nil, err
		}

		CreateActionRecord("Repair Record Create", "POST", "Repair Record", record, "Success")

		return result, nil
	} */

}

func GetOneRepairRecord(id string) (interface{}, error) {
	collection := config.GetCollection("repair_records")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var record models.RepairRecords

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid ID format")
	}

	filter := bson.M{ "_id": objectID }

	err = collection.FindOne(ctx, filter).Decode(&record)
	if err != nil {
		return nil, err
	}

	if record.Status == 0 {
		return "Oooops! This record has been removed!", nil
	} else {
		return &record, nil
	}
}

func UpdateRepairRecord(record *models.RepairRecords) (interface{}, error) {
	collection := config.GetCollection("repair_records")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// record.UpdatedAt = time.Now()

	filter := bson.M{"_id": record.ID}

	err := collection.FindOne(ctx, filter).Decode(&record)
	if err != nil {
		return nil, err
	}

	if record.Status == 0 {
		CreateActionRecord("Repair Record Update", "POST", "Repair Record", record, "Failed")
		return "Oooops! This record has been removed!", nil
	} else {
		record.UpdatedAt = time.Now()
		_, err := collection.ReplaceOne(ctx, filter, record)
		if err != nil {
			return nil, err	
		}
		CreateActionRecord("Repair Record Update", "POST", "Repair Record", record, "Success")
		return record, nil
	}
}

func VoidRepairRecord(id string) (interface{}, error) {
	collection := config.GetCollection("repair_records")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var record models.RepairRecords

	err = collection.FindOne(ctx, filter).Decode(&record)

	if err != nil {
		return nil, err
	}

	if record.Status == 1 {
		record.Status = 0
		record.UpdatedAt = time.Now()

		result, err := collection.UpdateOne(ctx, filter, bson.M{
			"$set": record,
		})

		if err != nil {
			return nil, err
		}

		CreateActionRecord("Repair Record Void", "DELETE", "Repair Record", record, "Success")
		return result, nil

	} else {
		CreateActionRecord("Repair Record Void", "DELETE", "Repair Record", record, "Failed")
		return "This record is already voided", nil
	}
}

func ListRepairRecords(dataReq *dto.RepairRecordPageReqDTO) (interface{}, error) {
	if dataReq.Page < 1 {
		dataReq.Page = 1
	}

	if dataReq.Limit < 1 {
		dataReq.Limit = 10
	}

	collection := config.GetCollection("repair_records")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	finalFilter := bson.M{"status": 1}
	if len(dataReq.DateRange) == 2 {
		finalFilter["createdAt"] = bson.M{
			"$gte": dataReq.DateRange[0],
			"$lte": dataReq.DateRange[1],
		}
	}

	assetMatch := bson.M{
		"$expr": bson.M{"$eq": []interface{}{"$_id", "$$assetIdStr"}},
	}

	if dataReq.AssetCode != "" {
		assetMatch["assetCode"] = dataReq.AssetCode
	}

	if len(dataReq.TypeIds) > 0 {
		objectTypeIds := make([]primitive.ObjectID, len(dataReq.TypeIds))
		for i, t := range dataReq.TypeIds {
			id, _ := primitive.ObjectIDFromHex(t)
			objectTypeIds[i] = id
		}
		assetMatch["typeId"] = bson.M{"$in": objectTypeIds}
	}

	if len(dataReq.DeptIds) > 0 {
		objectDeptIds := make([]primitive.ObjectID, len(dataReq.DeptIds))
		for i, d := range dataReq.DeptIds {
			id, _ := primitive.ObjectIDFromHex(d)
			objectDeptIds[i] = id
		}
		assetMatch["deptId"] = bson.M{"$in": objectDeptIds}
	}

	if len(dataReq.PlaceIds) > 0 {
		objectPlaceIds := make([]primitive.ObjectID, len(dataReq.PlaceIds))
		for i, p := range dataReq.PlaceIds {
			id, _ := primitive.ObjectIDFromHex(p)
			objectPlaceIds[i] = id
		}
		assetMatch["placeId"] = bson.M{"$in": objectPlaceIds}
	}

	skip := int64((dataReq.Page - 1) * dataReq.Limit)
    limit := int64(dataReq.Limit)

	pipeline := mongo.Pipeline{
    	bson.D{{Key: "$match", Value: finalFilter}},
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "asset_lists"},
				{Key: "let", Value: bson.D{
					{Key: "assetIdStr", Value: bson.D{
						{Key: "$convert", Value: bson.D{
							{Key: "input", Value: "$assetId"},
							{Key: "to", Value: "objectId"},
							{Key: "onError", Value: nil},
							{Key: "onNull", Value: nil},
						}},
					}},
				}},
				{Key: "pipeline", Value: mongo.Pipeline{
					bson.D{{Key: "$match", Value: assetMatch}},
					bson.D{{Key: "$unwind", Value: bson.M{"path": "$location", "preserveNullAndEmptyArrays": true}}},
				}},
				{Key: "as", Value: "assetlist"},
			}}},
			bson.D{{Key: "$unwind", Value: bson.M{"path": "$assetlist", "preserveNullAndEmptyArrays": true}}},
			bson.D{{Key: "$skip", Value: skip}},
			bson.D{{Key: "$limit", Value: limit}},
	}


	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var lists []bson.M
	if err := cursor.All(ctx, &lists); err != nil {
		return nil, err
	}

	countPipeline := mongo.Pipeline{
		{{Key: "$match", Value: finalFilter}},
		{{Key: "$lookup", Value: bson.M{
			"from": "assetlists",
			"let": bson.M{"assetIdStr": bson.M{"$toObjectId": "$assetId"}},
			"pipeline": mongo.Pipeline{
				{{Key: "$match", Value: assetMatch}},
			},
			"as": "assetlist",
		}}},
		{{Key: "$unwind", Value: bson.M{"path": "$assetlist", "preserveNullAndEmptyArrays": true}}},
		{{Key: "$count", Value: "total"}},
	}

	cursorTotal, err := collection.Aggregate(ctx, countPipeline)
	if err != nil {
		return nil, err
	}

	var totalRes []bson.M
	if err := cursorTotal.All(ctx, &totalRes); err != nil {
		return nil, err
	}

	total := int64(0)
	if len(totalRes) > 0 {
		if t, ok := totalRes[0]["total"].(int32); ok {
			total = int64(t)
		} else if t, ok := totalRes[0]["total"].(int64); ok {
			total = t
		}
	}

	return gin.H{"lists": lists, "total": total, "page": dataReq.Page, "limit": dataReq.Limit}, nil
}