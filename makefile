build:
	cd ./src && GOOS=linux go build -v -o ../bin/arbotgo
	cd ./src && GOOS=windows go build -v -o ../bin/arbotgo.exe
docker-build:
	docker build -t amaraa44/arbotgo:$(tag) -t amaraa44/arbotgo:latest .
docker-push:
	docker image push -a amaraa44/arbotgo