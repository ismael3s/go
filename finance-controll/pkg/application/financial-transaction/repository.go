package financialTransaction

type FinancialTransactionRepository interface {
	Create(financialTransaction *FinancialTransaction) error
	FindAll() ([]FinancialTransaction, error)
}
