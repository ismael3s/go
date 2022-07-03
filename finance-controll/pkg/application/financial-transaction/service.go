package financialTransaction

import "errors"

type FinancialTransactionService struct {
	repository FinancialTransactionRepository
}

func NewFinancialTransactionService(repository FinancialTransactionRepository) *FinancialTransactionService {
	return &FinancialTransactionService{
		repository: repository,
	}
}

func (s *FinancialTransactionService) Create(financialTransaction *FinancialTransaction) error {
	if financialTransaction.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	err := s.repository.Create(financialTransaction)

	if err != nil {
		return err
	}

	return nil
}

func (s *FinancialTransactionService) FindAll() ([]FinancialTransaction, error) {
	return s.repository.FindAll()
}
