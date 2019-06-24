BUILD := env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o

.PHONY: clean 

build: install 
	$(BUILD) ./bin/aws-subnet-exporter .

clean:
	rm -rf ./bin ./vendor

install:
	dep ensure -v