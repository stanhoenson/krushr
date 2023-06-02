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
	go test -cover ./...

.PHONY: coverage
coverage:
	@echo '1. Generating cover profile:'
	go test -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out

.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api
