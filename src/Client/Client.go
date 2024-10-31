package Client

import (
	"fmt"
	"net/rpc"
)

func Call(serverAddress string, serverPort int, callback func(*rpc.Client) error) {
	// Tenta uma conexão TCP com o banco
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, serverPort))
	if err != nil {
		fmt.Println("Client.Call : Failed to connect to Server : Error=", err)
		return
	}
	// Esse cara garante que a conexão será fechada graciosamente depois
	// de executar o callback. No melhor dos casos a função client.Close()
	// será chamada no final do escopo devido a palavra reservada "defer"
	defer client.Close()

	// Executa uma função "callback" que foi passada como parâmetro.
	err = callback(client)
	if err != nil {
		fmt.Println("Client.Call : Callback execution failed : Error=", err)
	}
}
