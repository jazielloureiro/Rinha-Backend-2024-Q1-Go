package entity

type Statement struct {
	Id          int    `json:"-"`
	AccountId   int    `json:"-"`
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
	Date        string `json:"realizada_em"`
}
