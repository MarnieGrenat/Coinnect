a:
	rm -f 	./Server ./Client
	go build ./src/Server/Server.go
	go build ./src/Client/Client.go

t:
	go test ./src/Server/Bank > test_log.txt

c:
	rm -f 	./Server ./Client