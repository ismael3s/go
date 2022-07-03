package financialTransaction

import "gorm.io/gorm"

type FinancialTransaction struct {
	gorm.Model
	ID     string `gorm:"primaryKey;unique"`
	Flow   string
	Amount float64
}
