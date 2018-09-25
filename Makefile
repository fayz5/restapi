PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test:
	@go test $(PKGS)
	echo $(DB_URL)

.PHONY: ensure
ensure:
	@dep ensure

.PHONY: clean
clean:
	@rm -rf $(CURDIR)/bin


