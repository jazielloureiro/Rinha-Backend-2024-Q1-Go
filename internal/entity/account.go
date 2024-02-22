package entity

type Account struct {
	Id    int `json:"-"`
	Limit int `json:"limite"`
	Value int `json:"saldo"`
}
