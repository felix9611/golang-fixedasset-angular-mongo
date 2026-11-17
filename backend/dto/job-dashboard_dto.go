package dto

type ListJobDashboardRequestDto struct {
	Name     string   `json:"name"`
	MemberId []string `json:"memberId"`
}
