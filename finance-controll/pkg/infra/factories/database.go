package factories

import (
	gormAdapter "github.com/ismael3s/go/finance-controll/pkg/infra/database/gorm_adapter"
	financialTransactionGorm "github.com/ismael3s/go/finance-controll/pkg/infra/database/gorm_adapter/financial_transaction"
	"gorm.io/gorm"
)

func NewGormDatabase(dsn string) (db *gorm.DB) {
	db = gormAdapter.NewGormDB(dsn)
	schemas := []interface{}{&financialTransactionGorm.FinancialTransaction{}}
	gormAdapter.RunMigrations(db, schemas...)

	return
}

func NewFinancialTransactionRepository(db *gorm.DB) *financialTransactionGorm.FinancialTransactionRepository {
	return financialTransactionGorm.NewFinancialTransactionRepository(db)
}
