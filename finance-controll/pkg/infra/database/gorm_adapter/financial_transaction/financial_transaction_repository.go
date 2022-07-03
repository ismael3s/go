package financialTransaction

import (
	teste "github.com/ismael3s/go/finance-controll/pkg/application/financial-transaction"
	"gorm.io/gorm"
)

var _ teste.FinancialTransactionRepository = new(FinancialTransactionRepository)

type FinancialTransactionRepository struct {
	db *gorm.DB
}

func NewFinancialTransactionRepository(db *gorm.DB) *FinancialTransactionRepository {
	return &FinancialTransactionRepository{db: db}
}

func (f *FinancialTransactionRepository) Create(financialTransaction *teste.FinancialTransaction) error {
	db := f.db.Create(&financialTransaction)

	return db.Error
}

func (f *FinancialTransactionRepository) FindAll() ([]teste.FinancialTransaction, error) {
	var financialTransactions []teste.FinancialTransaction

	db := f.db.Find(&financialTransactions)

	return financialTransactions, db.Error
}
