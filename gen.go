package gen

//go:generate go run ./cmd/tools/terndotenv/main.go run go generate
//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yaml
