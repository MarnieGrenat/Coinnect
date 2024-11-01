package BankManager

import (
	"fmt"
	"sync"
)

// account representa uma conta bancária com nome, senha e saldo.
type account struct {
	Name     string       // Nome do titular da conta.
	Password string       // Senha da conta.
	Balance  float64      // Saldo atual da conta.
	mutex    sync.RWMutex // RWMutex para operações seguras em nível de conta.
}

// Estruturas auxiliares para requests de login e operações.
type LoginRequest struct {
	AccountID       string
	AccountPassword string
}

type OperationRequest struct {
	AccountID       string
	AccountPassword string
	Quantity        float64
}

// Bank representa um banco que gerencia múltiplas contas.
type Bank struct {
	accounts map[int64]*account // Mapeia cada ID de conta para os dados da conta.
	nextID   int64              // ID da próxima conta a ser criada.
	mutex    sync.RWMutex       // RWMutex para garantir segurança em operações concorrentes em nível de banco.
}

// Initialize configura o banco inicializando o map de contas e criando uma conta de teste.
func (b *Bank) Initialize() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.accounts = make(map[int64]*account)
	b.nextID = 1

	// Conta hardcoded para teste
	b.accounts[b.nextID] = &account{
		Name:     "n",
		Password: "p",
		Balance:  2000,
	}
	b.nextID++
}

// OpenAccount cria uma nova conta bancária com o nome e a senha fornecidos.
func (b *Bank) OpenAccount(accountName string, accountPassword string, result *int64) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.accounts[b.nextID] = &account{
		Name:     accountName,
		Password: accountPassword,
		Balance:  0,
	}

	*result = b.nextID
	b.nextID++
	return nil
}

// CloseAccount fecha a conta com o ID fornecido, após autenticação da senha.
func (b *Bank) CloseAccount(accountID int64, accountPassword string, result *bool) error {
	_, isAuthenticated := b.getAuthenticatedAccount(accountID, accountPassword)
	if isAuthenticated {
		b.mutex.Lock()
		defer b.mutex.Unlock()

		delete(b.accounts, accountID)
		*result = true
		return nil
	}
	*result = false
	return fmt.Errorf("BankManager.CloseAccount : Failed to authenticate account : AccountID=%d", accountID)
}

// Withdraw realiza um saque de uma conta especificada, caso haja saldo suficiente.
func (b *Bank) Withdraw(accountID int64, accountPassword string, quantity float64, result *bool) error {
	account, isAuthenticated := b.getAuthenticatedAccount(accountID, accountPassword)

	if isAuthenticated {
		account.mutex.Lock()
		defer account.mutex.Unlock()

		if account.Balance >= quantity {
			account.Balance -= quantity
			*result = true
			return nil
		}
		*result = false
		return fmt.Errorf("BankManager.Withdraw : Insufficient funds for account : AccountID=%d", accountID)
	}
	*result = false
	return fmt.Errorf("BankManager.Withdraw : Failed to authenticate account : AccountID=%d", accountID)
}

// Deposit adiciona um valor ao saldo de uma conta especificada.
func (b *Bank) Deposit(accountID int64, accountPassword string, quantity float64, result *bool) error {
	account, isAuthenticated := b.getAuthenticatedAccount(accountID, accountPassword)

	if isAuthenticated {
		account.mutex.Lock()
		defer account.mutex.Unlock()

		account.Balance += quantity
		*result = true
		return nil
	}
	*result = false
	return fmt.Errorf("BankManager.Deposit : Failed to authenticate account : AccountID=%d", accountID)
}

// PeekBalance consulta o saldo de uma conta, se a senha estiver correta.
func (b *Bank) PeekBalance(accountID int64, accountPassword string, result *float64) error {
	account, isAuthenticated := b.getAuthenticatedAccount(accountID, accountPassword)

	if isAuthenticated {
		account.mutex.RLock()
		defer account.mutex.RUnlock()

		*result = account.Balance
		return nil
	}
	*result = 0
	return fmt.Errorf("BankManager.PeekBalance : Failed to authenticate account : AccountID=%d", accountID)
}

func (b *Bank) getAuthenticatedAccount(accountID int64, accountPassword string) (*account, bool) {
	account, accountExists := b.getAccount(accountID)
	if accountExists && (account.Password == accountPassword) {
		return account, true
	}
	return nil, false
}

// getAccount retorna uma conta segura e se ela existe.
func (b *Bank) getAccount(accountID int64) (*account, bool) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	info, exists := b.accounts[accountID]
	return info, exists
}
