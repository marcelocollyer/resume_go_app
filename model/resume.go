package model

// Resume struct (Model)
type Resume struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	CareerDesc  string       `json:"carrer_desc"`
	RolesDesc   string       `json:"roles_desc"`
	Experiences []Experience `json:"experiences"`
}
