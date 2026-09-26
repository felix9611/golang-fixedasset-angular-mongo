package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListStockTakeDto struct {
	PlaceIds []string `json:"placeIds"`
	Name     string   `json:"name"`
	Page     int64    `json:"page"`
	Limit    int64    `json:"limit"`
}
type GetStockTakeItemResponse struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StockTakeId primitive.ObjectID `bson:"stockTakeId" json:"stockTakeId"`
	AssetId     primitive.ObjectID `bson:"assetId" json:"assetId"`
	AssetCode   string             `bson:"assetCode" json:"assetCode"`
	PlaceId     string             `bson:"placeId" json:"placeId"`
	Status      string             `bson:"status" json:"status"`
	CheckTime   time.Time          `bson:"checkTime,omitempty" json:"checkTime"`
	Remark      string             `bson:"remark" json:"remark"`
}
