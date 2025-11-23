# Contributing to F5 Distributed Cloud Terraform Provider

Thank you for your interest in contributing to this project!

## Development Setup

### Prerequisites

- Go 1.21 or later
- Terraform 1.0 or later
- Make (optional, for using Makefile targets)
- F5 XC tenant with API credentials (for acceptance tests)

### Building

Clone the repository and build:

```bash
git clone https://github.com/robinmordasiewicz/terraform-provider-f5distributedcloud.git
cd terraform-provider-f5distributedcloud
go build ./...
```

### Testing

Run unit tests:

```bash
go test ./...
```

Run acceptance tests (requires F5 XC credentials):

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
TF_ACC=1 go test ./... -v
```

### Local Development

To use a locally built provider:

1. Build the provider:
   ```bash
   go build -o terraform-provider-f5distributedcloud
   ```

2. Create or update `~/.terraformrc`:
   ```hcl
   provider_installation {
     dev_overrides {
       "robinmordasiewicz/f5distributedcloud" = "/path/to/provider/binary"
     }
     direct {}
   }
   ```

3. Run Terraform commands without `terraform init`.

## Code Guidelines

### Style

- Follow Go conventions and idioms
- Use `gofmt` for formatting
- Run `go vet` to catch issues
- Use meaningful variable and function names

### Structure

- Resources go in `internal/resources/<resource_name>/`
- Data sources go in `internal/datasources/<datasource_name>/`
- Each resource/data source should have:
  - `model.go` - Data structures and API conversion
  - `schema.go` - Terraform schema definition
  - `resource.go` or `datasource.go` - CRUD operations
  - `resource_test.go` or `datasource_test.go` - Tests

### Documentation

- Update resource documentation in `docs/`
- Add examples in `examples/`
- Update CHANGELOG.md for notable changes

## Pull Request Process

1. Fork the repository
2. Create a feature branch from `main`
3. Make your changes with appropriate tests
4. Run all tests locally
5. Submit a pull request with a clear description

### PR Checklist

- [ ] Tests pass locally
- [ ] Code is formatted with `gofmt`
- [ ] Documentation is updated
- [ ] CHANGELOG.md is updated
- [ ] Commit messages are clear and descriptive

## Adding a New Resource

1. Create the resource directory:
   ```bash
   mkdir -p internal/resources/new_resource
   ```

2. Create the required files:
   - `model.go` - Define data model and API conversion
   - `schema.go` - Define Terraform schema
   - `resource.go` - Implement CRUD operations
   - `resource_test.go` - Write acceptance tests

3. Register the resource in `internal/provider/provider.go`

4. Add examples in `examples/resources/f5distributedcloud_new_resource/`

5. Add documentation in `docs/resources/new_resource.md`

## Adding a New Data Source

Similar to resources, but:
- Files go in `internal/datasources/new_datasource/`
- Only Read operation is needed
- Register in the `DataSources` function

## Reporting Issues

When reporting issues, please include:
- Terraform version
- Provider version
- Relevant configuration (sanitized)
- Expected vs actual behavior
- Error messages

## Code of Conduct

Please be respectful and constructive in all interactions.

## License

By contributing, you agree that your contributions will be licensed under the MPL-2.0 license.
