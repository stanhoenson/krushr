# API
.PHONY: run/api
run/api:
	go run ./cmd/api -debug -automigrate

.PHONY: build/api
build/api:
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api

.PHONY: audit/api
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

.PHONY: coverage/api
coverage:
	go test -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out

# UI
.PHONY: run/ui
run/ui:
	cd ui && \
	npm run start

.PHONY: build/ui
build/ui:
	cd ui && \
	npm run build

.PHONY: coverage/ui
coverage/ui:
	cd ui && \
	npm run coverage
