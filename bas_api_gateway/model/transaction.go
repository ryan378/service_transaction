package model

import "time"

type Transaction struct {
	ID              int    `gorm:"primarykey"`
	AccountID       string `gorm:"foreignkey"`
	BankID          string `gorm:"foreignkey"`
	Amount          int    `gorm:"column:amount"`
	TransactionDate *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
