TARGET := iskep

.PHONY: $(TARGET)

$(TARGET):
	go build -o $@ main.go

test:
	# brew install golangci-lint
	golangci-lint run ./...
