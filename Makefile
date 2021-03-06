mod:
	cd src && go mod download
	cd src && go mod verify
test:
	cd src && go test
build: mod test
	echo "package constants\n\nconst Version = \"$(TAG)\"" > src/constants/version.go
	cd src && GOOS=linux GOARCH=amd64 go build -o ../bin/arbotgo-linux
	cd src && GOOS=windows GOARCH=amd64 go build -o ../bin/arbotgo.exe
	cd src && GOOS=darwin GOARCH=amd64 go build -o ../bin/arbotgo-darwin