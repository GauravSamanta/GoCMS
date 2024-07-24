# Define variables
GOCMD=go
# TEMPL=templ
BUILD_DIR=./tmp
BINARY_NAME=main.exe

all: build


build: 
	$(GOCMD) build -v -o $(BUILD_DIR)/$(BINARY_NAME) .
	
clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

.PHONY: all build test clean