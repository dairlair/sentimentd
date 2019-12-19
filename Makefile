.PHONY: build
build:
	@echo "  >  Building binary..."
	go build -o build/sentimentd

.PHONY: migrate
migrate: build
	@echo "  >  Migrate..."
	build/sentimentd migrate