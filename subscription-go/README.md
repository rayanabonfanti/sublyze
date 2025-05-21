go mod tidy

go run main.go

go test -v ./handlers/...

# 1. Gerar cobertura para todos os pacotes
go test -coverprofile=coverage.out ./...

# 2. Gerar HTML
go tool cover -html=coverage.out -o coverage.html