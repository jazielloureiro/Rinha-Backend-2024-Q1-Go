package entity

type Statement struct {
	Id          int    `json:"id"`
	Value       int    `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
