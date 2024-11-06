package main

import (
	"Coinnect-FPPD/src/Client/Menu"
	Pygmalion "Coinnect-FPPD/src/deps"
	"fmt"
	"net/rpc"
	"strings"
	"time"

	"github.com/google/uuid"
)

// main is the entry point of the client application. It initializes the client
// configuration, reads the server address and port from the configuration file,
// and continuously performs client operations by sending requests to the server.
// If a request fails, it retries up to three times with exponential backoff.
func main() {
	// Carrega configurações
	Pygmalion.InitConfigReader("settings.yml", ".")
	port := Pygmalion.ReadInteger("ServerPort")
	address := Pygmalion.ReadString("ServerAddr")
	fmt.Printf("Client.main : Initializing Client : ServerAddress=%s, ServerPort=%d\n", address, port)

	for {
		operate(address, port, 3)
	}
}

// operate executes an operation by sending a request to a server at the specified address and port.
// It generates a unique request ID and obtains a callback function for the client operation.
// The function attempts to send the operation up to maxTries times, with exponential backoff between retries.
// If the error received contains "BankManager", the retries are stopped.
// Parameters:
//   - address: The server address to send the operation to.
//   - port: The server port to send the operation to.
func operate(address string, port int, maxTries int) {

	requestID := uuid.New().ID()
	callback := Menu.ObtainClientOperation(requestID)

	if callback != nil {
		// Executa a chamada ao servidor
		waitTime := 5 * time.Second
		for tries := 1; tries <= maxTries; tries++ {
			err := SendOperation(address, port, callback)

			// Verifica se a operação foi feita com sucesso
			if err == nil {
				break
			}

			// Verifica se o erro recebido pertence ao BankManager (erro tratado)
			if strings.Contains(err.Error(), "BankManager") {
				break
			}

			fmt.Printf("Client.SendOperation [Tries=%d] : Callback execution failed : Error=%s\n", tries, err)
			time.Sleep(waitTime)
			waitTime *= 2

			if tries == maxTries { // remove warning
				fmt.Println("Client.main : All retry attempts failed.")
			}
		}
	}
}

// SendOperation attempts to establish a TCP connection to the specified server
// address and port, and then executes a provided callback function with the
// established RPC client.
//
// Parameters:
//   - serverAddress: The address of the server to connect to.
//   - serverPort: The port of the server to connect to.
//   - callback: A function that takes an *rpc.Client and returns an error. This
//     function will be executed with the established RPC client.
//
// Returns:
//   - error: An error if the connection to the server fails or if the callback
//     function returns an error.
func SendOperation(serverAddress string, serverPort int, callback func(*rpc.Client) error) error {
	// Tenta uma conexão TCP com o banco
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, serverPort))
	if err != nil {
		fmt.Println("Client.SendOperation : Failed to connect to Server : Error=", err)
		return err
	}

	defer client.Close()

	// Executa uma função "callback" que foi passada como parâmetro.
	return callback(client)
}
