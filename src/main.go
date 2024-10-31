package main
import (
	Pygmalion "./deps/Pygmalion.go"
)

func main() {
	Pygmalion.InitConfigReader("settings.yml", ".")

	Addr := Pygmalion.ReadString("ServerAddr")
	Port := Pygmalion.ReadString("ServerPort")

	Server.Initialize(Addr, Port)
	fmt.Printf("Token: %s\n", Whitelist)
}