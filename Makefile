run-exchange-rate:
	swag init -g cmd/exchange-rate/main.go
	go run cmd/exchange-rate/main.go

run-test:
	go test -v ./...