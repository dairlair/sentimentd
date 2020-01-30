VERSION ?= $(shell git describe --tags 2>/dev/null | cut -c 2-)

.PHONY: build
build: clean
	@echo "  >  Building binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o ./build/sentimentd.linux-amd64 -ldflags='-X main.Version=$(VERSION) -extldflags "-static"' .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=vendor -a -o ./build/sentimentd.darwin-amd64 -ldflags='-X main.Version=$(VERSION) -extldflags "-static"' .
	cd ./build && find . -name 'sentimentd.*' | xargs -I{} tar czf {}.tar.gz {}
	cd ./build && shasum -a 256 * > sha256sum.txt
	cat ./build/sha256sum.txt

clean:
	@echo "  >  Cleaning up..."
	rm -rf build/*

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