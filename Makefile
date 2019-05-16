PROJECT_PKG = github.com/SteveLasker/dim
CLI_EXE     = dim
CLI_PKG     = $(PROJECT_PKG)/cmd/dim
LDFLAGS = -w
.PHONY: build
build: build-linux build-mac build-windows

.PHONY: build-linux
build-linux:
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -v --ldflags="$(LDFLAGS)" \
		-o bin/linux/amd64/$(CLI_EXE) $(CLI_PKG)

.PHONY: build-mac
build-mac:
	GOARCH=amd64 CGO_ENABLED=0 GOOS=darwin go build -v --ldflags="$(LDFLAGS)" \
		-o bin/darwin/amd64/$(CLI_EXE) $(CLI_PKG)

.PHONY: build-windows
build-windows:
	GOARCH=amd64 CGO_ENABLED=0 GOOS=windows go build -v --ldflags="$(LDFLAGS)" \
		-o bin/windows/amd64/$(CLI_EXE).exe $(CLI_PKG)