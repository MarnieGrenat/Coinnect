package BankBranch

import (
	"fmt"
	"net/rpc"
)

type OpenAccountRequest struct {
	Name      string
	Password  string
	RequestID int64
}

type AccountAccessRequest struct {
	AccountID int
	Password  string
	RequestID int64
}

type FundsOperationRequest struct {
	AccountID int
	Password  string
	Quantity  float64
	RequestID int64
}

func OpenNewAccount(name string, password string, requestID int64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := OpenAccountRequest{name, password, requestID}

		var response int

		err := client.Call("Bank.OpenAccount", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.OpenNewAccount : Server response=%d", response)
		return nil
	}
}

func CloseAccount(id int, password string, requestID int64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := AccountAccessRequest{id, password, requestID}

		var response bool

		err := client.Call("Bank.CloseAccount", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.CloseAccount : Server response=%t\n", response)
		return nil
	}
}

func Withdraw(id int, password string, quantity float64, requestID int64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := FundsOperationRequest{id, password, quantity, requestID}

		var response bool

		err := client.Call("Bank.Withdraw", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.Withdraw : Server response=%t\n", response)
		return nil
	}
}

func Deposit(id int, password string, quantity float64, requestID int64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := FundsOperationRequest{id, password, quantity, requestID}

		var response bool

		err := client.Call("Bank.Deposit", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.Deposit : Server response=%t\n", response)
		return nil
	}
}

func CheckBalance(id int, password string, requestID int64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := AccountAccessRequest{
			AccountID: id,
			Password:  password,
			RequestID: requestID,
		}
		fmt.Printf("Debug: Enviando Request - AccountID=%d, Password=%s, RequestID=%d\n", request.AccountID, request.Password, request.RequestID)

		var response float64

		err := client.Call("Bank.PeekBalance", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.CheckBalance : Server response=%.2f\n", response)
		return nil
	}
}
