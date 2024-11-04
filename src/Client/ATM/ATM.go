package ATM

import (
	"fmt"
	"net/rpc"
)

type AccountAccessRequest struct {
	AccountID int
	Password  string
	RequstID  int64
}

type FundsOperationRequest struct {
	AccountID int
	Password  string
	Quantity  float64
	RequestID int64
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
		request := AccountAccessRequest{id, password, requestID}

		var response float64

		err := client.Call("Bank.PeekBalance", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.CheckBalance : Server response=%.2f\n", response)
		return nil
	}
}
