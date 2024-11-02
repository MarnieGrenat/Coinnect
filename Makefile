all:
	go build ./src/Server/Server.go
	go build ./src/Client/Client.go

clean:
	rm -f 	./Server ./Client

