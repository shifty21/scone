GOCMD=go
GOBUILD=${GOCMD} build -compiler gccgo -compiler gccgo -a -installsuffix cgo
GOCLEAN=$(GOCMD) clean
BINARY_NAME=vault_initializer
    
all: clean build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v .

clean:
	$(GOCLEAN)
	rm -r $(BINARY_NAME)