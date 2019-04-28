GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOBIN=$(GOPATH)/bin
ZIPBUILD=$(GOBIN)/build-lambda-zip.exe
NOTIFIER=notifierlambda


build-all-linux: 
	$(MAKE) build-linux BINARY_NAME=$(NOTIFIER)

build: check-binary
	echo $(BINARY_NAME)
	GO111MODULE=on $(GOBUILD) -v ./cmd/aws/$(BINARY_NAME)/$(BINARY_NAME).go

build-linux: check-binary
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -v ./cmd/aws/$(BINARY_NAME)/$(BINARY_NAME).go
	
ifeq ("$(wildcard $(ZIPBUILD))","")
	GO111MODULE=on  $(GOGET) -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
endif

	$(GOBIN)/build-lambda-zip.exe -o $(BINARY_NAME).zip $(BINARY_NAME)


run: check-binary build
	./$(BINARY_NAME).exe

test:
	GO111MODULE=on $(GOTEST) -v ./...

build-notifier:
	$(MAKE) build BINARY_NAME=notifierlambda

clean: 
	rm *.exe *.zip notifierlambda

check-binary:
ifndef BINARY_NAME
	$(error BINARY_NAME is undefined)
endif


	