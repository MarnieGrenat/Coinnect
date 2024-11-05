package ATM

import (
	"fmt"
	"net/rpc"
)

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

func Withdraw(id int, password string, quantity float64, requestID uint32) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		request := FundsOperationRequest{id, password, quantity, requestID}

		var response bool

		err := client.Call("Bank.Withdraw", request, &response)
		if err != nil {
			return err
		}
		fmt.Printf("ATM.Withdraw : Operation has succeeded : RequestID=%d : ClientID=%d : HasSucceed=%t\n", requestID, request.AccountID, response)
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
		fmt.Printf("ATM.Deposit : Operation has succeeded : RequestID=%d : ClientID=%d : HasSucceed=%t\n", requestID, request.AccountID, response)
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
		fmt.Printf("ATM.CheckBalance : Checking Balance : RequestID=%d : ClientID=%d : Balance=%.2f\n", requestID, request.AccountID, response)
		return nil
	}
}
