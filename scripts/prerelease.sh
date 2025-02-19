GOOS=darwin GOARCH=amd64 go build -o bin/cli-macos ./cli/main.go
GOOS=linux GOARCH=amd64 go build -o bin/cli-linux ./cli/main.go
GOOS=windows GOARCH=amd64 go build -o bin/cli-windows.exe ./cli/main.go