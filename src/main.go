package main

import (
	"Coinnect-FPPD/src/Client"
	"Coinnect-FPPD/src/Client/ATM"
	"Coinnect-FPPD/src/Client/BankBranch"
	"Coinnect-FPPD/src/Server"
	Pygmalion "Coinnect-FPPD/src/deps"
)

func main() {
	// Carrega configurações
	Pygmalion.InitConfigReader("settings.yml", ".")
	port := Pygmalion.ReadInteger("ServerPort")
	address := Pygmalion.ReadString("ServerAddr")

	// Inicia o servidor
	go Server.Run(port)

	// Executa uma operação
	clientOpenAccountCallback := BankBranch.OpenNewAccount("Gabriela", "senhasegura")
	Client.Call(address, port, clientOpenAccountCallback)

	clientCheckBalanceCallback := ATM.CheckBalance(2, "senhasegura")
	Client.Call(address, port, clientCheckBalanceCallback)
}
