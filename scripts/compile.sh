#!/bin/bash

# Compile for Linux amd64
export GOOS=linux
export GOARCH=amd64
go build -o bin/envject-linux-amd64 main.go

# Compile for Linux 386
export GOOS=linux
export GOARCH=386
go build -o bin/envject-linux-386 main.go

# Compile for Linux arm64
export GOOS=linux
export GOARCH=arm64
go build -o bin/envject-linux-arm64 main.go

# Compile for Linux arm
export GOOS=linux
export GOARCH=arm
go build -o bin/envject-linux-arm main.go

# Compile for macOS amd64
export GOOS=darwin
export GOARCH=amd64
go build -o bin/envject-macos-amd64 main.go

# Compile for Macos arm64
export GOOS=darwin
export GOARCH=arm64
go build -o bin/envject-macos-arm64 main.go

# Compile for Windows amd64
export GOOS=windows
export GOARCH=amd64
go build -o bin/envject-windows-amd64.exe main.go

# Compile for Windows 386
export GOOS=windows
export GOARCH=386
go build -o bin/envject-windows-386.exe main.go

zip -r release.zip bin