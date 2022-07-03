package main

import (
	"log"
	"sync"

	financialTransaction "github.com/ismael3s/go/finance-controll/pkg/application/financial-transaction"
	"github.com/ismael3s/go/finance-controll/pkg/infra/factories"
)

var wg sync.WaitGroup

func main() {
	db := factories.NewGormDatabase("host=localhost user=root password=root dbname=root port=5435 sslmode=disable")
	repo := factories.NewFinancialTransactionRepository(db)

	service := financialTransaction.NewFinancialTransactionService(repo)

	transactions, _ := service.FindAll()

	for _, transaction := range transactions {
		wg.Add(1)
		go asyncGo(transaction)
	}

	wg.Wait()
}

func asyncGo(trasaction financialTransaction.FinancialTransaction) {
	defer wg.Done()
	log.Printf("%+v", trasaction)
}
