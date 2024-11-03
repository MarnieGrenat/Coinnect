a:
	go build ./src/Server/Server.go
	go build ./src/Client/Client.go

t:
	go test ./src/Server/Bank
	go test ./src/Client/ATM
	go test ./src/Client/BankBranch

c:
	rm -f 	./Server ./Client

