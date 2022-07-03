package financialTransaction

import "github.com/google/uuid"

type FinancialTransactionFlow string

const (
	Income  FinancialTransactionFlow = "income"
	Expense FinancialTransactionFlow = "expense"
)

type FinancialTransaction struct {
	ID     string                   `json:"id"`
	Flow   FinancialTransactionFlow `json:"flow"`
	Amount float64                  `json:"amount"`
}

func NewFinancialTransaction(flow FinancialTransactionFlow, amount float64) *FinancialTransaction {
	return &FinancialTransaction{
		ID:     uuid.New().String(),
		Flow:   flow,
		Amount: amount,
	}
}
