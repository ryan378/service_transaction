package model

type Bank struct {
	BankCode string `gorm:"primarykey"`
	Name     string
	Address  string
}

func (a *Bank) TableName() string {
	return "bank"
}
