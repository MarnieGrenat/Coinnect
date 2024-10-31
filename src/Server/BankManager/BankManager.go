package BankManager

type Account struct {
	Name string
	Password string // Tudo que eu fizer com psw são péssimas praticas mas não nos importamos com segurança aqui :)
	Balance float64
}

type AccountMap

type Bank struct {
	accounts map[int64]*Account // { AccountID : Account }
	nextID int64
}


func (b* Bank) Initialize() {
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

func (b* Bank) OpenAccount(accountName string, accountPassword string) bool{
	if isCredentialsValid(accountName, accountPassword) {
		b.accounts[.nextID] = &Account{
			Name: accountName,
			Password: accountPassword,
			Balance: 0,
		}
		b.nextID++
	}
}

func (b* Bank) CloseAccount(accountID int64, accountPassword string) bool{
	if isAuthenticated(accountID, accountPassword) {
		accounts.delete(accountID)
	}
}

func (b* Bank) Withdraw(accountID int64, accountPassword string, quantity float64) bool{
	if isAuthenticated(accountID, accountPassword) {
		account, _ := b.accounts
		if account.Balance >= quantity {
			account.Balance = account.Balance - quantity
			return true
		}
		return false
	}
	return false

}

func (b* Bank) Deposit(accountID int64 accountPassword string, quantity float64) bool{
	if isAuthenticated(accountID, accountPassword) {
		account, _ := b.accounts
		account.Balance = account.Balance + quantity
		return true
	}
	return false
}

func (b* Bank) PeekBalance(accountID int64, accountPassword string) float64{
	if isAuthenticated(accountID, accountPassword) {
		account, _ := b.accounts[accountID]
		return account.Balance
	}
	return 0
}

func (b *Bank) isAuthenticated(accountID int64, accountPassword string) bool{
	if accountID > 0 {
		accountInfo, accountExists := b.accounts[accountID]
		return accountExists && accountInfo.Password == accountPassword
	}
	return false
}

func (b *Bank) isCredentialsValid(accountID int64, accountPassword string) bool{
	return true
}
