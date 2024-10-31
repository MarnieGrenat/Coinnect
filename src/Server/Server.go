package Server

import (
	BankManager "Coinnect-FPPD/src/Server/Bank"
	"fmt"
	"net"
	"net/rpc"
)

func Run(port int) {
	bank := new(BankManager.Bank)
	bank.Initialize()

	// ainda não temos objetos thread-safe
	// e fugimos da idempotência. Um passo de cada vez 💪
	rpc.Register(bank)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Server.Run : Failed to initialize Server : Error=", err)
		return
	}

	for {
		fmt.Println("Server.Run : LocalHost at=", port)
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Server.Run : Failed to accept connection : Error=:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
