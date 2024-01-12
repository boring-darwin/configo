.PHONY: check
check:
	errcheck ./...
	go vet ./...

.PHONY: test
test:
	go test