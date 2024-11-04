package Menu

import (
	"Coinnect-FPPD/src/Client/ATM"
	"Coinnect-FPPD/src/Client/BankBranch"
	"fmt"
	"net/rpc"
	"strconv"
)

// Constantes para melhorar a leitura dos switch cases
type MenuChoice int

const (
	Exit MenuChoice = iota
	GoToATM
	GoToBankBranch
)

type OperationChoice int

const (
	Return OperationChoice = iota
	CheckBalance
	Deposit
	Withdraw
	OpenAccount
	CloseAccount
)

// Menu principal
func ObtainClientOperation(requestID int64) func(*rpc.Client) error {
	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1. Utilizar ATM")
		fmt.Println("2. Ir a um BankBranch")
		fmt.Println("3. Sair")

		var choice MenuChoice
		fmt.Print("Digite sua escolha: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch choice {
		case Exit:
			fmt.Println("Encerrando o programa.")
			return nil
		case GoToATM:
			operation := presentATMMenu(requestID)
			if operation != nil {
				return operation
			}
		case GoToBankBranch:
			operation := presentBankBranchMenu(requestID)
			// Vai retornar nil caso o cliente queira voltar ao menu principal
			if operation != nil {
				return operation
			}
		default:
			fmt.Println("Escolha inválida. Tente novamente")
		}
	}
}

func presentATMMenu(requestID int64) func(*rpc.Client) error {
	for {
		fmt.Println("\n--- Menu ATM ---")
		fmt.Println("0. Voltar ao menu principal")
		fmt.Println("1. Consultar saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Retirar")

		var choice OperationChoice
		fmt.Print("Digite sua escolha: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch choice {
		case Return:
			return nil
		case CheckBalance:
			id, password := getIDPasswordInput()
			return ATM.CheckBalance(id, password, requestID)
		case Deposit:
			id, password, amount := getIDPasswordAmountInput("depositar")
			return ATM.Deposit(id, password, amount, requestID)
		case Withdraw:
			id, password, amount := getIDPasswordAmountInput("retirar")
			return ATM.Withdraw(id, password, amount, requestID)
		default:
			fmt.Println("Escolha inválida. Tente novamente.")
		}
	}
}

func presentBankBranchMenu(requestID int64) func(*rpc.Client) error {
	for {
		fmt.Println("\n--- Menu BankBranch ---")
		fmt.Println("0. Voltar ao menu principal")
		fmt.Println("1. Consultar saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Retirar")
		fmt.Println("4. Abrir nova conta")
		fmt.Println("5. Fechar conta")

		var choice OperationChoice
		fmt.Print("Digite sua escolha: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch choice {
		case Return:
			return nil
		case CheckBalance:
			id, password := getIDPasswordInput()
			return BankBranch.CheckBalance(id, password, requestID)
		case Deposit:
			id, password, amount := getIDPasswordAmountInput("depositar")
			return BankBranch.Deposit(id, password, amount, requestID)
		case Withdraw:
			id, password, amount := getIDPasswordAmountInput("")
			return BankBranch.Withdraw(id, password, amount, requestID)
		case OpenAccount:
			name, password := getNamePasswordInput()
			return BankBranch.OpenNewAccount(name, password, requestID)
		case CloseAccount:
			id, password := getIDPasswordInput()
			return BankBranch.CloseAccount(id, password, requestID)
		default:
			fmt.Println("Escolha inválida. Tente novamente.")
		}
	}
}

func getNamePasswordInput() (string, string) {
	fmt.Print("Digite o nome da nova conta: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Digite a senha: ")
	var password string
	fmt.Scanln(&password)

	return name, password
}

func getIDPasswordInput() (int, string) {
	fmt.Print("Digite o ID da conta: ")
	var stringId string
	var password string

	_, err := fmt.Scanln(&stringId)
	if err != nil {
		fmt.Println("Erro ao ler o ID da conta. Certifique-se de inserir um número válido.")
	}

	fmt.Print("Digite a senha: ")
	fmt.Scanln(&password)

	id, _ := strconv.Atoi(stringId)
	fmt.Printf("ClientID: %d ClientPasswort:%s\n", id, password)
	return id, password
}

func getIDPasswordAmountInput(operation string) (int, string, float64) {
	id, password := getIDPasswordInput()

	fmt.Printf("Digite o valor a %s: ", operation)
	var amount float64
	fmt.Scanln(&amount)

	return id, password, amount
}
