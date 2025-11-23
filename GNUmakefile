default: build

HOSTNAME=registry.terraform.io
NAMESPACE=robinmordasiewicz
NAME=f5distributedcloud
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

# Generate provider schema for documentation
schema:
	@echo "Generating provider schema..."
	@rm -rf /tmp/tfschema && mkdir -p /tmp/tfschema
	@printf 'terraform {\n  required_providers {\n    f5distributedcloud = {\n      source = "robinmordasiewicz/f5distributedcloud"\n      version = "${VERSION}"\n    }\n  }\n}\n' > /tmp/tfschema/main.tf
	@cd /tmp/tfschema && terraform init -upgrade > /dev/null 2>&1
	@cd /tmp/tfschema && terraform providers schema -json | python3 -c "import json,sys;d=json.load(sys.stdin);k=list(d['provider_schemas'].keys())[0];d['provider_schemas']['f5distributedcloud']=d['provider_schemas'].pop(k);print(json.dumps(d))" > $(CURDIR)/schema.json
	@echo "Schema generated: schema.json"

# Generate documentation using pre-generated schema
docs: schema
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name f5distributedcloud --providers-schema schema.json

clean:
	rm -f ${BINARY} schema.json

.PHONY: build install test testacc lint fmt generate schema docs clean
