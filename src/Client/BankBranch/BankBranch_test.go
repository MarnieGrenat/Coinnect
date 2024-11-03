package BankBranch

import (
	"net/rpc"
	"testing"
)

func TestOpenNewAccount(t *testing.T) {
	client := &rpc.Client{}
	callback := OpenNewAccount("TestUser", "password")

	err := callback(client)
	if err != nil {
		t.Errorf("Expected no error when opening a new account, got %v", err)
	}
}

func TestCloseAccount(t *testing.T) {
	client := &rpc.Client{}
	callback := CloseAccount(1, "password")

	err := callback(client)
	if err != nil {
		t.Errorf("Expected no error when closing an account, got %v", err)
	}
}
