# Define variables
AIR_CMD := air
GO_CMD := go
TEMPL_CMD := templ

all: build

build:
	 $(TEMPL_CMD) generate
    # $(GO_CMD) build -o main .

clean:
    rm -f main

    