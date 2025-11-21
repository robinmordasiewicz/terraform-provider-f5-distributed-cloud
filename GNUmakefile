default: build

HOSTNAME=registry.terraform.io
NAMESPACE=robinmordasiewicz
NAME=f5-distributed-cloud
BINARY=terraform-provider-${NAME}
VERSION=0.1.0
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -v -cover ./...

testacc:
	TF_ACC=1 go test -v -cover ./... -timeout 120m

lint:
	golangci-lint run ./...

fmt:
	gofmt -s -w .

generate:
	go generate ./...

docs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

clean:
	rm -f ${BINARY}

.PHONY: build install test testacc lint fmt generate docs clean
