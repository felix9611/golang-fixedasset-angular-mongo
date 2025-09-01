package dto

type TaxInformationListDTO struct {
	NameCode   	string				`json:"nameCode"`
	Tax         string				`json:"tax"`
	Page       int64              `json:"page"`
	Limit      int64              `json:"limit"`
}
