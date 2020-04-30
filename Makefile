generate:
	go run internal/generator.go -template internal/providers.go.template > providers.go
	goimports -w .