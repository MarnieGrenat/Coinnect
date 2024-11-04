package main

import (
	BankManager "Coinnect-FPPD/src/Server/Bank"
	Pygmalion "Coinnect-FPPD/src/deps"
	"fmt"
	"net"
	"net/rpc"
)

var listener net.Listener

func main() {
	fmt.Printf("Server.main : Initializing Operations")
	// Carrega configurações
	Pygmalion.InitConfigReader("settings.yml", ".")
	port := Pygmalion.ReadInteger("ServerPort")

	// Inicia o servidor
	Run(port)
	// Garante que o servidor feche graciosamente :)
	defer Close()
	fmt.Printf("Server.main : Finishing Operations")
}

func Run(port int) {
	bank := new(BankManager.Bank)
	bank.Initialize()

	rpc.Register(bank)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Server.Run : Failed to initialize Server : Error=", err)
		return
	}

	fmt.Println("Server.Run : LocalHost at=", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Server.Run : Failed to accept connection : Error=:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

// Close encerra o listener e para de aceitar novas conexões
func Close() {
	if listener != nil {
		err := listener.Close()
		if err != nil {
			fmt.Println("Server.Close : Failed to close listener : Error=", err)
		} else {
			fmt.Println("Server.Close : Server closed successfully.")
		}
	} else {
		fmt.Println("Server.Close : No active listener found.")
	}
}
