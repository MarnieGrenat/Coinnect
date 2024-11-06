package BankManager

import (
	"fmt"
	"sync"
	"time"
)

// Estruturas auxiliares para requests de login e operações.
type OpenAccountRequest struct {
	Name      string
	Password  string
	RequestID uint32
}

type AccountAccessRequest struct {
	AccountID int
	Password  string
	RequestID uint32
}

type FundsOperationRequest struct {
	AccountID int
	Password  string
	Quantity  float64
	RequestID uint32
}

// account representa uma conta bancária com nome, senha e saldo.
type account struct {
	Name     string       // Nome do titular da conta.
	Password string       // Senha da conta.
	Balance  float64      // Saldo atual da conta.
	mutex    sync.RWMutex // RWMutex para operações seguras em nível de conta.
}

// Bank representa um banco que gerencia múltiplas contas.
type Bank struct {
	accounts          map[int]*account       // Mapeia cada ID de conta para os dados da conta.
	nextID            int                    // ID da próxima conta a ser criada.
	mutex             sync.RWMutex           // RWMutex para garantir segurança em operações concorrentes em nível de banco.
	processedRequests map[uint32]interface{} // Log de resultados de operações indexados por RequestID.
	requestLogMutex   sync.Mutex             // Mutex para controlar acesso ao log de operações.
}

// Initialize configura o banco inicializando o map de contas e criando uma conta de teste.
func (b *Bank) Initialize() {
	fmt.Println("BankManager.Initialize : Initializing bank.")
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.accounts = make(map[int]*account)
	b.nextID = 0
	b.processedRequests = make(map[uint32]interface{})

	// Conta hardcoded para teste
	b.accounts[b.nextID] = &account{
		Name:     "n",
		Password: "p",
		Balance:  2000,
	}
	b.nextID++
	fmt.Printf("BankManager.Initialize : Finished initializing : Next usable ClientID=%d\n", b.nextID)
}

// OpenAccount cria uma nova conta bancária com o nome e a senha fornecidos.
func (b *Bank) OpenAccount(request OpenAccountRequest, result *int) error {
	if previousResult, exists := b.checkRequestID(request.RequestID); exists {
		*result = previousResult.(int)
		return nil
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()
	fmt.Printf("BankManager.OpenAccount [RequestID=%d] : Opening a new account : AccountID=%d : Name=%s\n", request.RequestID, b.nextID, request.Name)
	b.accounts[b.nextID] = &account{
		Name:     request.Name,
		Password: request.Password,
		Balance:  0,
	}

	*result = b.nextID
	b.logRequestID(request.RequestID, *result)
	b.nextID++
	fmt.Printf("BankManager.OpenAccount [RequestID=%d] : Opened a new account successfully : Next Usable ID=%d\n", request.RequestID, b.nextID)
	return nil
}

// CloseAccount fecha a conta com o ID fornecido, após autenticação da senha.
func (b *Bank) CloseAccount(request AccountAccessRequest, result *bool) error {
	if previousResult, exists := b.checkRequestID(request.RequestID); exists {
		*result = previousResult.(bool)
		return nil
	}

	account, isAuthenticated := b.getAuthenticatedAccount(request.AccountID, request.Password)
	if !isAuthenticated {
		*result = false
		fmt.Printf("BankManager.CloseAccount [RequestID=%d] : Failed to authenticate account : AccountID=%d : AccountPassword=%s\n", request.RequestID, request.AccountID, request.Password)
		b.logRequestID(request.RequestID, *result)
		return fmt.Errorf("BankManager.CloseAccount [RequestID=%d] : Failed to authenticate account : AccountID=%d", request.RequestID, request.AccountID)
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	fmt.Printf("BankManager.CloseAccount [RequestID=%d] : Closing account : AccountID=%d : Balance=%.2f : ClientName=%s\n", request.RequestID, request.AccountID, account.Balance, account.Name)
	delete(b.accounts, request.AccountID)
	*result = true
	b.logRequestID(request.RequestID, *result)
	return nil
}

// Withdraw realiza um saque de uma conta especificada, caso haja saldo suficiente.
func (b *Bank) Withdraw(request FundsOperationRequest, result *bool) error {
	if previousResult, exists := b.checkRequestID(request.RequestID); exists {
		*result = previousResult.(bool)
		return nil
	}

	account, isAuthenticated := b.getAuthenticatedAccount(request.AccountID, request.Password)

	if !isAuthenticated {
		*result = false
		b.logRequestID(request.RequestID, *result)
		fmt.Printf("BankManager.Withdraw [RequestID=%d] : Failed to authenticate account : AccountID=%d\n", request.RequestID, request.AccountID)
		return fmt.Errorf("BankManager.Withdraw [RequestID=%d] : Failed to authenticate account : AccountID=%d", request.RequestID, request.AccountID)
	}

	account.mutex.Lock()
	defer account.mutex.Unlock()

	if account.Balance >= request.Quantity {
		account.Balance -= request.Quantity
		*result = true
		b.logRequestID(request.RequestID, *result)
		fmt.Printf("BankManager.Withdraw [RequestID=%d] : Withdrawing funds : AccountID=%d : Balance=%.2f : Quantity=%.2f\n", request.RequestID, request.AccountID, account.Balance, request.Quantity)
		return nil
	}
	*result = false
	b.logRequestID(request.RequestID, *result)
	fmt.Printf("BankManager.Withdraw [RequestID=%d] : Insufficient funds for account : AccountID=%d : Quantity=%.2f\n", request.RequestID, request.AccountID, request.Quantity)
	return fmt.Errorf("BankManager.Withdraw [RequestID=%d] : Insufficient funds for account : AccountID=%d", request.RequestID, request.AccountID)
}

// Deposit adiciona um valor ao saldo de uma conta especificada.
func (b *Bank) Deposit(request FundsOperationRequest, result *bool) error {
	if previousResult, exists := b.checkRequestID(request.RequestID); exists {
		*result = previousResult.(bool)
		return nil
	}

	account, isAuthenticated := b.getAuthenticatedAccount(request.AccountID, request.Password)

	if !isAuthenticated {
		*result = false
		b.logRequestID(request.RequestID, *result)
		fmt.Printf("BankManager.Deposit [RequestID=%d] : Failed to authenticate account : AccountID=%d\n", request.RequestID, request.AccountID)
		return fmt.Errorf("BankManager.Deposit [RequestID=%d] : Failed to authenticate account : AccountID=%d", request.RequestID, request.AccountID)
	}

	account.mutex.Lock()
	defer account.mutex.Unlock()

	account.Balance += request.Quantity
	*result = true
	b.logRequestID(request.RequestID, *result)
	fmt.Printf("BankManager.Deposit [RequestID=%d] : Depositing on account : AccountID=%d : Balance=%.2f : Quantity=%.2f\n", request.RequestID, request.AccountID, account.Balance, request.Quantity)
	return nil
}

// PeekBalance consulta o saldo de uma conta, se a senha estiver correta.
func (b *Bank) PeekBalance(request AccountAccessRequest, result *float64) error {
	if previousResult, exists := b.checkRequestID(request.RequestID); exists {
		*result = previousResult.(float64)
		return nil
	}
	time.Sleep(10 * time.Second)

	account, isAuthenticated := b.getAuthenticatedAccount(request.AccountID, request.Password)
	if !isAuthenticated {
		*result = -1
		b.logRequestID(request.RequestID, *result)
		fmt.Printf("BankManager.PeekBalance [RequestID=%d] : Failed to authenticate account : AccountID=%d : AccountPassword=%s\n", request.RequestID, request.AccountID, request.Password)
		return fmt.Errorf("BankManager.PeekBalance [RequestID=%d] : Failed to authenticate account : AccountID=%d", request.RequestID, request.AccountID)
	}
	account.mutex.RLock()
	defer account.mutex.RUnlock()

	*result = account.Balance
	b.logRequestID(request.RequestID, *result)

	fmt.Printf("BankManager.PeekBalance [RequestID=%d] : Peeking balance : AccountID=%d : Balance=%.2f\n", request.RequestID, request.AccountID, account.Balance)
	return nil
}

// Funções auxiliares

func (b *Bank) getAuthenticatedAccount(AccountID int, accountPassword string) (*account, bool) {
	account, accountExists := b.getAccount(AccountID)
	if accountExists && (account.Password == accountPassword) {
		return account, true
	}
	return nil, false
}

func (b *Bank) getAccount(AccountID int) (*account, bool) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	info, exists := b.accounts[AccountID]
	return info, exists
}

// checkRequestID verifica se uma requisição já foi processada.
func (b *Bank) checkRequestID(requestID uint32) (interface{}, bool) {
	b.requestLogMutex.Lock()
	defer b.requestLogMutex.Unlock()

	result, exists := b.processedRequests[requestID]
	if exists {
		fmt.Printf("BankManager.checkRequestID [RequestID=%d] : Request already replied\n", requestID)
	}
	return result, exists
}

// logRequestID registra o resultado de uma operação para um RequestID.
func (b *Bank) logRequestID(requestID uint32, result interface{}) {
	b.requestLogMutex.Lock()
	defer b.requestLogMutex.Unlock()

	b.processedRequests[requestID] = result
}
