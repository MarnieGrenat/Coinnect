package BankBranch

import (
	"fmt"
	"net/rpc"
)

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

func OpenNewAccount(name string, password string, requestID uint32) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := OpenAccountRequest{name, password, requestID}

		var response int

		err := client.Call("Bank.OpenAccount", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.OpenNewAccount : Account Created : RequestID=%d : ClientID=%d\n", requestID, response)
		return nil
	}
}

func CloseAccount(id int, password string, requestID uint32) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := AccountAccessRequest{id, password, requestID}

		var response bool

		err := client.Call("Bank.CloseAccount", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.CloseAccount : Account Closed Successfully : RequestID=%d : ClientID=%d : AccountClosed=%t\n", requestID, request.AccountID, response)
		return nil
	}
}

func Withdraw(id int, password string, quantity float64, requestID uint32) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := FundsOperationRequest{id, password, quantity, requestID}

		var response bool

		err := client.Call("Bank.Withdraw", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.Withdraw : Operation has succeeded : RequestID=%d : ClientID=%d : HasSucceed=%t\n", requestID, request.AccountID, response)
		return nil
	}
}

func Deposit(id int, password string, quantity float64, requestID uint32) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := FundsOperationRequest{id, password, quantity, requestID}

		var response bool

		err := client.Call("Bank.Deposit", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.Deposit : Operation has succeeded : RequestID=%d : ClientID=%d : HasSucceed=%t\n", requestID, request.AccountID, response)
		return nil
	}
}

func CheckBalance(id int, password string, requestID uint32) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := AccountAccessRequest{id, password, requestID}

		var response float64

		err := client.Call("Bank.PeekBalance", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("BankBranch.CheckBalance : Checking Balance : RequestID=%d : ClientID=%d : Balance=%.2f\n", requestID, request.AccountID, response)
		return nil
	}
}
