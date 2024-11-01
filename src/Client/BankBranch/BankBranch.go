package BankBranch

import (
	"fmt"
	"net/rpc"
)

type OpenAccountRequest struct {
	Name     string
	Password string
}

type AccountAccessRequest struct {
	ID       int64
	Password string
}

type FundsOperationRequest struct {
	ID       int64
	Password string
	Quantity float64
}

func OpenNewAccount(name string, password string) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := OpenAccountRequest{name, password}

		var response int64

		err := client.Call("Bank.OpenAccount", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.OpenNewAccount : Server response=%d", response)
		return nil
	}
}

func CloseAccount(id int64, password string) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := AccountAccessRequest{id, password}

		var response bool

		err := client.Call("Bank.CloseAccount", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.CloseAccount : Server response=%t\n", response)
		return nil
	}
}

func Withdraw(id int64, password string, quantity float64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := FundsOperationRequest{id, password, quantity}

		var response bool

		err := client.Call("Bank.Withdraw", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.Withdraw : Server response=%t\n", response)
		return nil
	}
}

func Deposit(id int64, password string, quantity float64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := FundsOperationRequest{id, password, quantity}

		var response bool

		err := client.Call("Bank.Deposit", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.Deposit : Server response=%t\n", response)
		return nil
	}
}

func CheckBalance(id int64, password string) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := AccountAccessRequest{id, password}

		var response float64

		err := client.Call("Bank.PeekBalance", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.CheckBalance : Server response=%.2f\n", response)
		return nil
	}
}
