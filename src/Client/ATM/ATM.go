package ATM

import (
	"fmt"
	"net/rpc"
)

func Withdraw(id int64, password string, quantity float64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		return withdraw(client, id, password, quantity)
	}
}

func withdraw(client *rpc.Client, id int64, password string, quantity float64) error {
	request := struct {
		ID       int64
		Password string
		Quantity float64
	}{id, password, quantity}

	var response bool

	err := client.Call("Bank.Withdraw", request, &response)
	if err != nil {
		return err
	}
	fmt.Printf("ATM.Withdraw : Server response=%t\n", response)
	return nil
}

func Deposit(id int64, password string, quantity float64) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		return deposit(client, id, password, quantity)
	}
}

func deposit(client *rpc.Client, id int64, password string, quantity float64) error {
	request := struct {
		ID       int64
		Password string
		Quantity float64
	}{id, password, quantity}

	var response bool

	err := client.Call("Bank.Deposit", request, &response)
	if err != nil {
		return err
	}
	fmt.Printf("ATM.Deposit : Server response=%t\n", response)
	return nil
}

func CheckBalance(id int64, password string) func(*rpc.Client) error {
	return func(client *rpc.Client) error {
		return checkBalance(client, id, password)
	}
}

func checkBalance(client *rpc.Client, id int64, password string) error {
	request := struct {
		ID       int64
		Password string
	}{id, password}

	var response float64

	err := client.Call("Bank.PeekBalance", request, &response)
	if err != nil {
		return err
	}
	fmt.Printf("ATM.CheckBalance : Server response=%.2f\n", response)
	return nil
}
