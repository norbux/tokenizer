default:
	@go build -o ./cmd/main.exe ./cmd/main.go
	@./cmd/main.exe -f ./cmd/input.txt
