.PHONY: audit
audit:
	@echo '1. Tidying and verifying module dependencies:'
	go mod tidy
	go mod verify
	@echo ''

	@echo '2. Formatting code:'
	go fmt ./...
	gofumpt -l -w .
	@echo ''

	@echo '3. Vetting code:'
	go vet ./...
	staticcheck ./...
	@echo ''

	@echo '4. Running tests:'
	go test -race -vet=off ./...

.PHONY: coverage
coverage:
	@echo '1. Generating cover profile:'
	go test -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out
