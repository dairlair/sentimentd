.PHONY: build
build:
	@echo "  >  Building binary..."
	go build -o build/sentimentd

test: build
	go test -short -coverprofile=bin/cov.out `go list ./... | grep -v vendor/`
	go tool cover -func=build/cov.out

.PHONY: migrate
migrate: build
	@echo "  >  Migrate..."
	build/sentimentd migrate