app_name := tracker
app_module := github.com/dstgo/tracker/cmd/$(app_name)
author := github.com/dstgo
version := $(shell git describe --tags --always)
git_version := $(shell git describe --tags --always)
build_time := $(shell date +"%Y%m%d%H%M%S")
host_os := $(shell go env GOHOSTOS)
host_arch := $(shell go env GOHOSTARCH)
go_os := $(shell go env GOOS)

# build target
target := $(app_name)

# windows platform
ifeq ($(go_os), windows)
	target := $(app_name).exe
endif

.PHONY: init
init:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/gin-swagger@latest
	go get github.com/swaggo/files@latest
	go install github.com/google/wire/cmd/wire@latest
	go get github.com/google/wire/cmd/wire@latest

.PHONY: gen
gen:
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: build
build:
	go vet ./...
	go build -trimpath \
				-ldflags="-X main.Author=$(author) -X main.Version=$(version) -X main.BuildTime=$(build_time)" \
				-o ./bin/$(app_name)/$(target) $(app_module)