package BankManager

import (
	"sync"
	"testing"

	"github.com/google/uuid"
)

const CONCURRENCY_QUANTITY int = 1000
const IDEMPONTENCY_QUANTITY int = 10000

func TestOpenAccount(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	request := OpenAccountRequest{Name: "TestUser", Password: "password", RequestID: uuid.New().ID()}
	var accountID int

	err := bank.OpenAccount(request, &accountID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if accountID != 1 { // Verifica se a próxima conta é criada com ID correto
		t.Errorf("Expected account ID to be 1, got %d", accountID)
	}
}

func TestCloseAccount_Success(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	// Criar uma conta para ser fechada
	request := OpenAccountRequest{Name: "TestUser", Password: "password", RequestID: uuid.New().ID()}
	var accountID int
	_ = bank.OpenAccount(request, &accountID)

	closeRequest := AccountAccessRequest{AccountID: accountID, Password: "password", RequestID: uuid.New().ID()}
	var result bool

	err := bank.CloseAccount(closeRequest, &result)
	if err != nil || !result {
		t.Fatalf("Expected account to close successfully, got error %v", err)
	}

	// Verificar se a conta realmente foi removida
	var balance float64
	balanceRequest := AccountAccessRequest{AccountID: accountID, Password: "password", RequestID: uuid.New().ID()}
	err = bank.PeekBalance(balanceRequest, &balance)
	if err == nil {
		t.Errorf("Expected error when peeking balance of closed account, but got none")
	}
}

func TestWithdraw_Concurrency(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	request := OpenAccountRequest{Name: "ConcurrentUser", Password: "password", RequestID: uuid.New().ID()}
	var accountID int
	_ = bank.OpenAccount(request, &accountID)

	var totalFunds float64 = 1000000000000
	var withdrawFunds float64 = 1

	depositRequest := FundsOperationRequest{AccountID: accountID, Password: "password", Quantity: totalFunds, RequestID: uuid.New().ID()}
	var depositResult bool
	_ = bank.Deposit(depositRequest, &depositResult)

	var wg sync.WaitGroup
	numWithdraws := CONCURRENCY_QUANTITY
	for i := 0; i < numWithdraws; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			withdrawRequest := FundsOperationRequest{AccountID: accountID, Password: "password", Quantity: withdrawFunds, RequestID: uuid.New().ID()}
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
	balanceRequest := AccountAccessRequest{AccountID: accountID, Password: "password", RequestID: uuid.New().ID()}
	_ = bank.PeekBalance(balanceRequest, &balance)
	expectedBalance := totalFunds - float64(numWithdraws*int(withdrawFunds))
	if balance != expectedBalance {
		t.Errorf("Expected balance to be %.2f, got %.2f", expectedBalance, balance)
	}
}

func TestDeposit_Idempotency(t *testing.T) {
	bank := &Bank{}
	bank.Initialize()

	request := OpenAccountRequest{Name: "IdempotentUser", Password: "password", RequestID: uuid.New().ID()}
	var accountID int
	_ = bank.OpenAccount(request, &accountID)

	depositRequest := FundsOperationRequest{AccountID: accountID, Password: "password", Quantity: 500, RequestID: uuid.New().ID()}
	var depositResult bool

	// Realiza múltiplos depósitos com o mesmo RequestID para testar idempotência
	for i := 0; i < IDEMPONTENCY_QUANTITY; i++ {
		err := bank.Deposit(depositRequest, &depositResult)
		if err != nil {
			t.Errorf("Deposit failed with error: %v", err)
		}
	}

	var balance float64
	balanceRequest := AccountAccessRequest{AccountID: accountID, Password: "password", RequestID: uuid.New().ID()}
	_ = bank.PeekBalance(balanceRequest, &balance)

	if balance != 500 {
		t.Errorf("Expected balance to be 500 after idempotent deposits, got %.2f", balance)
	}
}
