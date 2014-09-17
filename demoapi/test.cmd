cd %GOPATH%\src\github.com\agoravoting\agora-http-go
goose up
godep go test github.com/agoravoting/agora-http-go/demoapi
goose down