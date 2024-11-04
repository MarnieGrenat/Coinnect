a:
	go build ./src/Server/Server.go
	go build ./src/Client/Client.go

t:
	go test ./src/Server/Bank

c:
	rm -f 	./Server ./Client

