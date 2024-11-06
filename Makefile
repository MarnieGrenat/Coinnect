a:
	rm -f 	./Server ./Client
	go build ./src/Server/Server.go
	go build ./src/Client/Client.go

t:
	go test ./src/Server/Bank
	go test ./src/Client/

c:
	rm -f 	./Server ./Client