package ATM

import (
	"errors"
	"net/rpc"
	"testing"
)

func mockClientCall(method string, args interface{}, reply interface{}) error {
	if method == "Bank.PeekBalance" {
		*reply.(*float64) = 500.0
		return nil
	}
	return errors.New("method not implemented")
}

func TestCheckBalance(t *testing.T) {
	client := &rpc.Client{}
	callback := CheckBalance(1, "password")

	err := callback(client)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
