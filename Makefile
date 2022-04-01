build:
	cd src && go mod download
	cd src && go mod verify

	cd src && GOOS=linux GOARCH=amd64 go build -o ../bin/arbotgo-linux
	cd src && GOOS=windows GOARCH=amd64 go build -o ../bin/arbotogo.exe
	cd src && GOOS=darwin GOARCH=amd64 go build -o ../bin/arbotogo-darwin
docker-build:
	docker build -t amaraa44/arbotgo:$(tag) -t amaraa44/arbotgo:latest .
docker-push:
	docker image push -a amaraa44/arbotgo