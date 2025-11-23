# Terraform Provider for F5 Distributed Cloud

[![Registry](https://img.shields.io/badge/terraform-registry-blue)](https://registry.terraform.io/providers/robinmordasiewicz/f5-distributed-cloud)
[![Go Report Card](https://goreportcard.com/badge/github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud)](https://goreportcard.com/report/github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud)
[![License: MPL-2.0](https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg)](LICENSE)

Terraform provider for managing F5 Distributed Cloud (F5 XC) resources.

## Overview

This provider enables infrastructure engineers to manage F5 Distributed Cloud resources using Terraform, including:

- **Namespaces** - Organizational units for resources
- **Origin Pools** - Backend server pools
- **HTTP Load Balancers** - Layer 7 load balancing
- **Application Firewalls** - WAF policies
- **Cloud Sites** - AWS VPC, Azure VNET, and GCP VPC sites
- **Cloud Credentials** - Authentication for cloud providers
- And many more resources...

## Installation

### From Terraform Registry (Recommended)

```hcl
terraform {
  required_providers {
    f5distributedcloud = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 0.1"
    }
  }
}
```

### Manual Installation

Download the appropriate binary from the [releases page](https://github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/releases) and place it in:

- **Linux/macOS**: `~/.terraform.d/plugins/`
- **Windows**: `%APPDATA%\terraform.d\plugins\`

## Quick Start

### 1. Configure Authentication

Set environment variables:

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
```

Or use a P12 certificate:

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_P12_FILE="/path/to/certificate.p12"
export F5XC_API_P12_PASSWORD="your-password"
```

### 2. Configure the Provider

```hcl
provider "f5distributedcloud" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.api_token
}
```

### 3. Create Resources

```hcl
# Create a namespace
resource "f5distributedcloud_namespace" "example" {
  name        = "my-namespace"
  description = "Example namespace"
}

# Create an origin pool
resource "f5distributedcloud_origin_pool" "backend" {
  name      = "backend-pool"
  namespace = f5distributedcloud_namespace.example.name
  port      = 8080

  origin_servers {
    public_name {
      dns_name = "backend.example.com"
    }
  }
}

# Create an HTTP load balancer
resource "f5distributedcloud_http_loadbalancer" "frontend" {
  name      = "frontend-lb"
  namespace = f5distributedcloud_namespace.example.name
  domains   = ["app.example.com"]
  http_port = 80

  default_pool {
    name      = f5distributedcloud_origin_pool.backend.name
    namespace = f5distributedcloud_namespace.example.name
  }
}
```

## Documentation

- [Provider Documentation](docs/index.md)
- [Resource Documentation](docs/resources/)
- [Data Source Documentation](docs/data-sources/)
- [Examples](examples/)
- [Guides](docs/guides/)

## Resources

| Resource | Description |
|----------|-------------|
| `f5distributedcloud_namespace` | Manage namespaces |
| `f5distributedcloud_origin_pool` | Manage origin pools |
| `f5distributedcloud_http_loadbalancer` | Manage HTTP load balancers |
| `f5distributedcloud_tcp_loadbalancer` | Manage TCP load balancers |
| `f5distributedcloud_app_firewall` | Manage application firewalls |
| `f5distributedcloud_service_policy` | Manage service policies |
| `f5distributedcloud_cloud_credentials` | Manage cloud credentials |
| `f5distributedcloud_aws_vpc_site` | Manage AWS VPC sites |
| `f5distributedcloud_azure_vnet_site` | Manage Azure VNET sites |
| `f5distributedcloud_gcp_vpc_site` | Manage GCP VPC sites |
| `f5distributedcloud_voltstack_site` | Manage Voltstack sites |

See [docs/resources/](docs/resources/) for full documentation.

## Data Sources

| Data Source | Description |
|-------------|-------------|
| `f5distributedcloud_namespace` | Read namespace data |
| `f5distributedcloud_origin_pool` | Read origin pool data |
| `f5distributedcloud_http_loadbalancer` | Read HTTP load balancer data |
| `f5distributedcloud_tcp_loadbalancer` | Read TCP load balancer data |
| `f5distributedcloud_app_firewall` | Read application firewall data |
| `f5distributedcloud_cloud_credentials` | Read cloud credentials data |
| `f5distributedcloud_virtual_site` | Read virtual site data |
| `f5distributedcloud_service_policy` | Read service policy data |
| `f5distributedcloud_healthcheck` | Read healthcheck configuration |
| `f5distributedcloud_rate_limiter` | Read rate limiter configuration |

See [docs/data-sources/](docs/data-sources/) for full documentation.

## Building from Source

```bash
# Clone the repository
git clone https://github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud.git
cd terraform-provider-f5-distributed-cloud

# Build the provider
go build -o terraform-provider-f5xc

# Install locally
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/robinmordasiewicz/f5-distributed-cloud/0.1.0/$(go env GOOS)_$(go env GOARCH)
mv terraform-provider-f5xc ~/.terraform.d/plugins/registry.terraform.io/robinmordasiewicz/f5-distributed-cloud/0.1.0/$(go env GOOS)_$(go env GOARCH)/
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history.

## License

This project is licensed under the Mozilla Public License 2.0 - see the [LICENSE](LICENSE) file for details.
