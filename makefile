build:
	GOOS=linux go build -o ./bin/arbotgo
	GOOS=windows go build -o ./bin/arbotgo.exe