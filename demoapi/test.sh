#!/bin/bash
cd $GOPATH/src/github.com/agoravoting/agora-http-go
goose up
go test -v github.com/agoravoting/agora-http-go/agora-http-go
goose down