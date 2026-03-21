package dto

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"time"
)

type ListAssetReqDto struct {
	Page          int      `json:"page"`
	Limit         int      `json:"limit"`
	AssetCode     string   `json:"assetCode"`
	AssetName     string   `json:"assetName"`
	TypeIds       []string `json:"typeIds"`
	PlaceIds      []string `json:"placeIds"`
	DeptIds       []string `json:"deptIds"`
	PurchaseDates []string `json:"purchaseDates"`
}

type ListRecordReqDto struct {
	Page      int      `json:"page"`
	Limit     int      `json:"limit"`
	AssetCode string   `json:"assetCode"`
	DateRange []string `json:"dateRange"`
}

type AssetListList struct {
	Lists []models.AssetLists `json:"lists"`
	Total int64               `json:"total"`
	Page  int                 `json:"page"`
	Limit int                 `json:"limit"`
}

type DashboardReqFilterDto struct {
	TypeIds       []string    `json:"typeIds"`
	PlaceIds      []string    `json:"placeIds"`
	DeptIds       []string    `json:"deptIds"`
	PurchaseDates []time.Time `json:"purchaseDates"`
}

type DashboardReqDto struct {
	DataType      bool                   `json:"dataType"`
	DataTypeValue string                 `json:"dataTypeValue"`
	DateType      bool                   `json:"dateType"`
	DateTypeValue string                 `json:"dateTypeValue"`
	ValueField    string                 `json:"valueField"`
	Filter        *DashboardReqFilterDto `json:"filter"`
}
