package main

import (
	"Coinnect-FPPD/src/Client"
	"Coinnect-FPPD/src/Menu"
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
	// Garante que o servidor feche graciosamente :)
	defer Server.Close()

	// Executa uma operação
	callback := Menu.ObtainClientOperation()

	// Executa a chamada ao servidor
	Client.SendOperation(address, port, callback)
}
