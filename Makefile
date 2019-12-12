.PHONY: build
build:
	@echo "  >  Building binary..."
	go build -o build/

.PHONY: train
train: build
	@echo "  >  Train with dataset..."
	# Create the test brain
	export brain_id=`build/sentimentd brain create skynet | grep skynet | tail -n 1 | awk {'print $1'}`
	cat dataset.txt | build/sentimentd train $(brain_id)