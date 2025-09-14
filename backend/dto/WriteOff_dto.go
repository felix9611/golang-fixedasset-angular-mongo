package dto

import (
	"time"
)

type CreateWriteOffRecrod struct {
	AssetId        string  `json:"assetId"`
	LastPlaceId    string  `json:"lastPlaceId"`
	Reason         string  `json:"reason"`
	LastDay        string  `json:"lastDay"`
	DisposalMethod string  `json:"disposalMethod"`
	RemainingValue float64 `json:"remainingValue"`
}

type ListWriteOffReqDto struct {
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	PlaceIds  []string    `json:"placeIds"`
	DeptIds   []string    `json:"deptIds"`
	TypeIds   []string    `json:"typeIds"`
	DateRange []time.Time `json:"dateRange"`
}
