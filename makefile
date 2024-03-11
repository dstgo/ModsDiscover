app_package := github.com/dstgo/tracker/cmd/tracker
app_name := tracker
go_os := $(shell go env GOOS)
target := $(app_name)

ifeq ($(go_os), windows)
	target := $(app_name).exe
endif

.PHONY: build
build:
	go vet ./...
	go build -trimpath -o ./bin/$(app_name)/$(target) $(app_package)
