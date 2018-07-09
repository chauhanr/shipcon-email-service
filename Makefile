build:
	GOOS=linux GOARCH=amd64 go build -o shipcon-email-service
	docker build -t shipcon-email-service .
	go clean

run:
	docker run -d --net="host" \
		-p 50054 \
		-e MICRO_SERVER_ADDRESS=:50054 \
		-e MICRO_REGISTRY=mdns \
        shipcon-email-service