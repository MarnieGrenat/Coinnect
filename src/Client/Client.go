package main

import (
	"Coinnect-FPPD/src/Client/Menu"
	Pygmalion "Coinnect-FPPD/src/deps"
	"fmt"
	"sync"
	"net/rpc"
)

var mutex sync.Mutex

func main() {
	// Carrega configurações
	Pygmalion.InitConfigReader("settings.yml", ".")
	port := Pygmalion.ReadInteger("ServerPort")
	address := Pygmalion.ReadString("ServerAddr")
	fmt.Printf("Client.main : Initializing Client : ServerAddress=%s, ServerPort=%d\n", address, port)
	var requestID int64 = 1
	for {
		// Executa uma operação
		callback := Menu.ObtainClientOperation(requestID)
		if callback != nil {
			// Executa a chamada ao servidor
			mutex.Lock()
			SendOperation(address, port, callback)
			requestID++
			mutex.Unlock()
			continue
		}
		break
	}
}

func SendOperation(serverAddress string, serverPort int, callback func(*rpc.Client) error) {
	// Tenta uma conexão TCP com o banco
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, serverPort))
	if err != nil {
		fmt.Println("Client.SendOperation : Failed to connect to Server : Error=", err)
		return
	}
	// Esse cara garante que a conexão será fechada graciosamente depois
	// de executar o callback. No melhor dos casos a função client.Close()
	// será chamada no final do escopo devido a palavra reservada "defer"
	defer client.Close()

	// Executa uma função "callback" que foi passada como parâmetro.
	err = callback(client)
	if err != nil {
		fmt.Println("Client.SendOperation : Callback execution failed : Error=", err)
	}
}
