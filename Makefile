GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
BUILD_DIR=build
BUILD_NAME=arvi

.PHONY: build

all: test build
br: build run
build: 
		$(GOBUILD) -o $(BUILD_DIR)/$(BUILD_NAME)
test: 
		$(GOTEST) -v .
clean: 
		$(GOCLEAN)
		rm -r $(BUILD_DIR)
run:
		./$(BUILD_DIR)/$(BUILD_NAME)