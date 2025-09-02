package dto

type ListAssetReqDto struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	AssetCode string `json:"assetCode"`
	AssetName string `json:"assetName"`
	TypeIds   []string `json:"typeIds"`
	PlaceIds  []string `json:"placeIds"`
	DeptIds   []string `json:"deptIds"`
	PurchaseDates []string `json:"purchaseDates"`
}

type ListRecordReqDto struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	AssetCode string `json:"assetCode"`
	DateRange []string `json:"dateRange"`
}