package Menu

import (
	"Coinnect-FPPD/src/Client/ATM"
	"Coinnect-FPPD/src/Client/BankBranch"
	"fmt"
	"net/rpc"
	"os"
)

func ObtainClientOperation() func(*rpc.Client) error {

	// Menu principal
	for {
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1. Utilizar ATM")
		fmt.Println("2. Ir a um BankBranch")
		fmt.Println("3. Sair")

		var choice int
		fmt.Print("Digite sua escolha: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch choice {
		case 1:
			operation := menuATM()
			if operation != nil {
				return operation
			}
		case 2:
			operation := menuBankBranch()
			// Vai retornar nil caso o cliente queira voltar ao menu principal
			if operation != nil {
				return operation
			}

		case 3:
			fmt.Println("Encerrando o programa.")
			os.Exit(1)
		default:
			fmt.Println("Escolha inválida. Tente novamente.")
		}
	}
}

func menuATM() func(*rpc.Client) error {
	for {
		fmt.Println("\n--- Menu ATM ---")
		fmt.Println("1. Consultar saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Retirar")
		fmt.Println("4. Voltar ao menu principal")

		var choice int
		fmt.Print("Digite sua escolha: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Digite o ID da conta: ")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			return ATM.CheckBalance(id, password)
		case 2:
			fmt.Print("Digite o ID da conta: ")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			fmt.Print("Digite o valor a depositar: ")
			var amount float64
			fmt.Scanln(&amount)
			return ATM.Deposit(id, password, amount)
		case 3:
			fmt.Print("Digite o ID da conta: ")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			fmt.Print("Digite o valor a retirar: ")
			var amount float64
			fmt.Scanln(&amount)
			return ATM.Withdraw(id, password, amount)
		case 4:
			return nil
		default:
			fmt.Println("Escolha inválida. Tente novamente.")
		}
	}
}

func menuBankBranch() func(*rpc.Client) error {
	for {
		fmt.Println("\n--- Menu BankBranch ---")
		fmt.Println("1. Consultar saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Retirar")
		fmt.Println("4. Abrir nova conta")
		fmt.Println("5. Fechar conta")
		fmt.Println("6. Voltar ao menu principal")

		var choice int
		fmt.Print("Digite sua escolha: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Digite o ID da conta: ")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			return BankBranch.CheckBalance(id, password)
		case 2:
			fmt.Print("Digite o ID da conta: ")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			fmt.Print("Digite o valor a depositar: ")
			var amount float64
			fmt.Scanln(&amount)
			return BankBranch.Deposit(id, password, amount)
		case 3:
			fmt.Print("Digite o ID da conta: ")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			fmt.Print("Digite o valor a retirar: ")
			var amount float64
			fmt.Scanln(&amount)
			return BankBranch.Withdraw(id, password, amount)
		case 4:
			fmt.Print("Digite o nome da nova conta: ")
			var name string
			fmt.Scanln(&name)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			return BankBranch.OpenNewAccount(name, password)
		case 5:
			fmt.Print("Digite o ID da conta: ")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("Digite a senha: ")
			var password string
			fmt.Scanln(&password)
			return BankBranch.CloseAccount(id, password)
		case 6:
			return nil
		default:
			fmt.Println("Escolha inválida. Tente novamente.")
		}
	}
}
