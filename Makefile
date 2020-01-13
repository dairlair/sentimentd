.PHONY: build
build:
	@echo "  >  Building binary..."
	go build -o build/sentimentd

test: build
	go test -short -coverprofile=build/cov.out `go list ./... | grep -v vendor/`
	go tool cover -func=build/cov.out

.PHONY: migrate
migrate: build
	@echo "  >  Migrate..."
	build/sentimentd migrate

.PHONY: bench
bench:
	@echo "  >  Benchmarking..."
	go test -bench=. ./...

.PHONY: mocks
mocks:
	@echo " > Generate mocks..."
	mockery -all -keeptree -dir pkg -output ./pkg/mocks