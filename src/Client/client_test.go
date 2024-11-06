package main

import (
	"errors"
	"net/rpc"
	"strings"
	"testing"
	"time"
)

// TestOperateSuccess testa o caso em que a operação é bem-sucedida na primeira tentativa.
func TestOperateSuccess(t *testing.T) {
	mockCallback := func(client *rpc.Client) error {
		return nil // simula sucesso imediato
	}
	operate("localhost", 8080, 3, mockCallback)
}

// TestOperateRetries testa o comportamento de retentativas ao falhar em estabelecer conexão.
func TestOperateRetries(t *testing.T) {
	failCallback := func(client *rpc.Client) error {
		return errors.New("mock connection error") // simula falha
	}

	start := time.Now()
	operate("localhost", 8080, 3, failCallback)
	duration := time.Since(start)

	expectedMinDuration := 5*time.Second + 10*time.Second + 20*time.Second // tempo de backoff esperado
	if duration < expectedMinDuration {
		t.Errorf("Expected duration to be at least %v, got %v", expectedMinDuration, duration)
	}
}

// TestOperateBankManagerError testa o caso em que o erro "BankManager" interrompe as retentativas.
func TestOperateBankManagerError(t *testing.T) {
	mockCallbackWithError := func(client *rpc.Client) error {
		return errors.New("BankManager: simulated error")
	}
	operate("localhost", 8080, 3, mockCallbackWithError)
}

// TestSendOperationSuccess verifica a conexão ao servidor e execução do callback com sucesso.
func TestSendOperationSuccess(t *testing.T) {
	mockCallback := func(client *rpc.Client) error {
		return nil // simula sucesso do callback
	}
	err := SendOperation("localhost", 8080, mockCallback)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// TestSendOperationConnectionFailure testa uma falha de conexão ao servidor.
func TestSendOperationConnectionFailure(t *testing.T) {
	mockCallback := func(client *rpc.Client) error {
		return nil
	}
	err := SendOperation("invalid_address", 8080, mockCallback)
	if err == nil || !strings.Contains(err.Error(), "Failed to connect to Server") {
		t.Errorf("Expected connection error, got %v", err)
	}
}

// mockCallback é uma função de callback de teste que simula uma execução de operação no servidor.
func mockCallback(client *rpc.Client) error {
	if client == nil {
		return errors.New("mock client error")
	}
	return nil
}

// mockCallbackWithError simula um erro específico de "BankManager".
func mockCallbackWithError(client *rpc.Client) error {
	return errors.New("BankManager: simulated error")
}
