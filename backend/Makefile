.PHONY: build clean deploy

build:
	# dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/manage-connection manage-connection/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/create-channel create-channel/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/list-channels list-channels/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/subscribe-channel subscribe-channel/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/send-message send-message/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/broadcast-message broadcast-message/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
