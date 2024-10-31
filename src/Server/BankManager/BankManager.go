package BankManager

import (
	"fmt"
)

// Account representa uma conta bancária com nome, senha e saldo.
type Account struct {
	Name     string  // Nome do titular da conta.
	Password string  // Senha da conta.
	Balance  float64 // Saldo atual da conta.
}

// Bank representa um banco que gerencia múltiplas contas usando um map para armazenar contas e um ID para controle.
type Bank struct {
	accounts map[int64]*Account // Mapeia cada ID de conta para os dados da conta.
	nextID   int64              // ID da próxima conta a ser criada.
}

// Initialize configura o banco inicializando o map de contas e criando uma conta de teste.
// Ele deve ser chamado antes de outras operações do banco.
func (b *Bank) Initialize() {
	b.accounts = make(map[int64]*Account)
	b.nextID = 1

	// Conta hardcoded para teste
	b.accounts[b.nextID] = &Account{
		Name:     "n",
		Password: "p",
		Balance:  2000,
	}
	b.nextID++
}

// OpenAccount cria uma nova conta bancária com o nome e a senha fornecidos.
// Parâmetros:
// - accountName: o nome do titular da nova conta.
// - accountPassword: a senha para a nova conta.
// - result: um ponteiro para um bool que será true se a conta for criada com sucesso.
// Retorna:
// - error: um erro, caso a conta não seja criada por algum motivo.
func (b *Bank) OpenAccount(accountName string, accountPassword string, result *bool) error {
	b.accounts[b.nextID] = &Account{
		Name:     accountName,
		Password: accountPassword,
		Balance:  0,
	}

	_, accountExists := b.accounts[b.nextID]
	if accountExists {
		b.nextID++
		*result = true
		return nil
	}
	return fmt.Errorf("BankManager.OpenAccount : Failed to create a new account : AccountName=%s", accountName)
}

// CloseAccount fecha a conta com o ID fornecido, após autenticação da senha.
// Parâmetros:
// - accountID: o ID da conta a ser fechada.
// - accountPassword: a senha da conta para autenticação.
// - result: um ponteiro para um bool que será true se a conta for fechada com sucesso.
// Retorna:
// - error: um erro caso a conta não seja fechada, como erro de autenticação.
func (b *Bank) CloseAccount(accountID int64, accountPassword string, result *bool) error {
	if b.isAuthenticated(accountID, accountPassword) {
		delete(b.accounts, accountID)
		_, accountExists := b.accounts[accountID]
		if !accountExists {
			*result = true
			return nil
		}
		return fmt.Errorf("BankManager.CloseAccount : Failed to delete account : AccountID=%d", accountID)
	}
	*result = false
	return fmt.Errorf("BankManager.CloseAccount : Failed to authenticate account : AccountID=%d", accountID)
}

// Withdraw realiza um saque de uma conta especificada, caso haja saldo suficiente.
// Parâmetros:
// - accountID: o ID da conta de onde o valor será retirado.
// - accountPassword: a senha da conta para autenticação.
// - quantity: o valor a ser sacado.
// - result: um ponteiro para um bool que será true se o saque for bem-sucedido.
// Retorna:
// - error: um erro caso o saque não seja realizado, como erro de autenticação ou saldo insuficiente.
func (b *Bank) Withdraw(accountID int64, accountPassword string, quantity float64, result *bool) error {
	if b.isAuthenticated(accountID, accountPassword) {
		account := b.accounts[accountID]
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
// Parâmetros:
// - accountID: o ID da conta que receberá o valor.
// - accountPassword: a senha da conta para autenticação.
// - quantity: o valor a ser depositado.
// - result: um ponteiro para um bool que será true se o depósito for bem-sucedido.
// Retorna:
// - error: um erro caso o depósito não seja realizado, como erro de autenticação.
func (b *Bank) Deposit(accountID int64, accountPassword string, quantity float64, result *bool) error {
	if b.isAuthenticated(accountID, accountPassword) {
		account := b.accounts[accountID]
		account.Balance += quantity
		*result = true
		return nil
	}
	*result = false
	return fmt.Errorf("BankManager.Deposit : Failed to authenticate account : AccountID=%d", accountID)
}

// PeekBalance consulta o saldo de uma conta, se a senha estiver correta.
// Parâmetros:
// - accountID: o ID da conta que será consultada.
// - accountPassword: a senha da conta para autenticação.
// - result: um ponteiro para float64 que armazenará o saldo da conta se a consulta for bem-sucedida.
// Retorna:
// - error: um erro caso a consulta não seja realizada, como erro de autenticação.
func (b *Bank) PeekBalance(accountID int64, accountPassword string, result *float64) error {
	if b.isAuthenticated(accountID, accountPassword) {
		account := b.accounts[accountID]
		*result = account.Balance
		return nil
	}
	*result = 0
	return fmt.Errorf("BankManager.PeekBalance : Failed to authenticate account : AccountID=%d", accountID)
}

// isAuthenticated verifica se a senha fornecida está correta para a conta especificada.
// Esta função é usada internamente e não deve ser chamada diretamente por outros pacotes.
// Parâmetros:
// - accountID: o ID da conta a ser autenticada.
// - accountPassword: a senha da conta a ser verificada.
// Retorna:
// - bool: true se a senha estiver correta, false caso contrário.
func (b *Bank) isAuthenticated(accountID int64, accountPassword string) bool {
	accountInfo, accountExists := b.accounts[accountID]
	return accountExists && (accountInfo.Password == accountPassword)
}
