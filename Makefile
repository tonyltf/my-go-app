run:
	swag init
	go run main.go

run-test:
	go test -v ./...