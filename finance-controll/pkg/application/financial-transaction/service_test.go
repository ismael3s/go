package financialTransaction

import "testing"

type financialTransactionRepositoryStub struct {
	finalTransactions []FinancialTransaction
}

func (f financialTransactionRepositoryStub) Create(financialTransaction *FinancialTransaction) error {
	t := NewFinancialTransaction(financialTransaction.Flow, financialTransaction.Amount)
	f.finalTransactions = append(f.finalTransactions, *t)
	return nil
}

func (f financialTransactionRepositoryStub) FindAll() ([]FinancialTransaction, error) {
	return f.finalTransactions, nil
}

func TestFinancialTransactionService_Create(t *testing.T) {
	type args struct {
		financialTransaction *FinancialTransaction
	}
	tests := []struct {
		name         string
		s            *FinancialTransactionService
		args         args
		errorMessage string
	}{
		{
			name: "should return error when amount is less than 0",
			s: NewFinancialTransactionService(financialTransactionRepositoryStub{
				[]FinancialTransaction{},
			}),
			errorMessage: "amount must be greater than 0",
			args: args{
				financialTransaction: &FinancialTransaction{
					ID:     "",
					Flow:   Income,
					Amount: -1,
				},
			},
		},
	}

	for _, tt := range tests {
		//

		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Create(tt.args.financialTransaction)

			if got.Error() != tt.errorMessage {
				t.Errorf("FinancialTransactionService.Create() = %v, want %v", got, tt.errorMessage)
			}
		})
	}
}
