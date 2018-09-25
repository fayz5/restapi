PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test:
	@go test $(PKGS)

.PHONY: ensure
ensure:
	@dep ensure


