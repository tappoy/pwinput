WORKING_DIRS=tmp bin
SRC=$(shell find -name "*.go")
BIN=bin/$(shell basename $(CURDIR))
DOC=Document.txt
COVER=tmp/cover
COVER0=tmp/cover0

.PHONY: all clean fmt cover test lint run

all: $(WORKING_DIRS) $(FMT) $(BIN) test lint $(DOC) $(TESTBIN)

clean:
	rm -rf $(WORKING_DIRS)

$(WORKING_DIRS):
	mkdir -p $(WORKING_DIRS)

fmt: $(SRC)
	go fmt ./... > $(FMT) 2>&1 || true

go.sum: go.mod
	go mod tidy

$(BIN): go.sum $(SRC)
	go build -o $(BIN) cmd/main.go

test: $(BIN)
	go test -v -tags=mock -cover -coverprofile=$(COVER) ./...

$(COVER0): $(COVER)
	grep "0$$" $(COVER) | tee > $(COVER0) 2>&1

cover: $(COVER)
	go tool cover -html=$(COVER)

$(DOC): $(SRC)
	go doc -all . > $(DOC)

lint: $(BIN)
	go vet
