
# the name of the binary when built
BINARY_NAME=rpsweb
MAIN_PACKAGE_PATH := ./main.go

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

# remove any binaries that are built
.PHONY: clean
clean:
	rm -f ./bin/$(BINARY_NAME)*

.PHONY: build-debug
build-debug: clean
	CGO_ENABLED=0 go build -gcflags=all="-N -l" -o bin/$(BINARY_NAME)-debug ${MAIN_PACKAGE_PATH}

.PHONY: run
run: build-debug
	./bin/$(BINARY_NAME)-debug