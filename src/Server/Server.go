package Server

import (
    "fmt"
    "net"
    "net/rpc"
    BankManager "./Server/BankManager/BankManager.go"
)


func RunBank(port str) {
    bankServer := new(Bank)
    bankServer.inicializar()

    rpc.Register(bankServer)
    l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        fmt.Println("Server : Failed to initialize Server : Error=", err)
        return
    }

    for {
        fmt.Println("Server : Localhost at=", port)
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Server : Failed to accept connection : Error=:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}