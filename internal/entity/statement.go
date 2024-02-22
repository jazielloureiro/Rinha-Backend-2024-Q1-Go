package entity

type Statement struct {
	Id          int    `json:"id"`
	AccountId   int    `json:"accountId"`
	Value       int    `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
