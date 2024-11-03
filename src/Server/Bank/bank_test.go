package BankManager

import (
	"sync"
	"testing"
)

func TestOpenAccount(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	request := OpenAccountRequest{Name: "TestUser", Password: "password"}
	var accountID int

	err := bank.OpenAccount(request, &accountID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if accountID != 1 {
		t.Errorf("Expected account ID to be 1, got %d", accountID)
	}
}

func TestCloseAccount_Success(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	// Criar uma conta para ser fechada
	request := OpenAccountRequest{Name: "TestUser", Password: "password"}
	var accountID int
	_ = bank.OpenAccount(request, &accountID)

	closeRequest := AccountAccessRequest{AccountID: accountID, Password: "password"}
	var result bool

	err := bank.CloseAccount(closeRequest, &result)
	if err != nil || !result {
		t.Fatalf("Expected account to close successfully, got error %v", err)
	}
}

func TestWithdraw_Concurrency(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	request := OpenAccountRequest{Name: "ConcurrentUser", Password: "password"}
	var accountID int
	_ = bank.OpenAccount(request, &accountID)

	depositRequest := FundsOperationRequest{AccountID: accountID, Password: "password", Quantity: 1000}
	var depositResult bool
	_ = bank.Deposit(depositRequest, &depositResult)

	var wg sync.WaitGroup
	numWithdraws := 100
	for i := 0; i < numWithdraws; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			withdrawRequest := FundsOperationRequest{AccountID: accountID, Password: "password", Quantity: 10}
			var withdrawResult bool
			err := bank.Withdraw(withdrawRequest, &withdrawResult)
			if err != nil {
				t.Errorf("Withdraw failed with error: %v", err)
			}
		}()
	}
	wg.Wait()

	// Verifica se o saldo está correto após as retiradas
	var balance float64
	balanceRequest := AccountAccessRequest{AccountID: accountID, Password: "password"}
	_ = bank.PeekBalance(balanceRequest, &balance)
	expectedBalance := 1000 - float64(numWithdraws*10)
	if balance != expectedBalance {
		t.Errorf("Expected balance to be %.2f, got %.2f", expectedBalance, balance)
	}
}

func TestDeposit_Idempotency(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	request := OpenAccountRequest{Name: "IdempotentUser", Password: "password"}
	var accountID int
	_ = bank.OpenAccount(request, &accountID)

	depositRequest := FundsOperationRequest{AccountID: accountID, Password: "password", Quantity: 500}
	var depositResult bool

	// Realiza múltiplos depósitos idênticos para testar idempotência
	for i := 0; i < 3; i++ {
		_ = bank.Deposit(depositRequest, &depositResult)
	}

	var balance float64
	balanceRequest := AccountAccessRequest{AccountID: accountID, Password: "password"}
	_ = bank.PeekBalance(balanceRequest, &balance)

	if balance != 500 {
		t.Errorf("Expected balance to be 500 after idempotent deposits, got %.2f", balance)
	}
}
