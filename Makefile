# Define variables
AIR_CMD := air
GO_CMD := go
TEMPL_CMD := templ

all: build

build:
	$(TEMPL_CMD) generate
    $(GO_CMD) build -o ./tmp/main.exe .

clean:
    rm -f main

.PHONY: all build test clean
    