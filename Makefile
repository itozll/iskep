TARGET := iskep

.PHONY: $(TARGET)

$(TARGET):
	go build -o $@ main.go
