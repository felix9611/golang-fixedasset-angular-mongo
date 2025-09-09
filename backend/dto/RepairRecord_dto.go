package dto

type RepairRecordPageReqDTO struct {
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
	DateRange []string `json:"dateRange"`
	AssetCode string `json:"assetCode"`
	DeptIds []string `json:"deptIds"`
	TypeIds []string `json:"typeIds"`
	PlaceIds []string `json:"placeIds"`
}