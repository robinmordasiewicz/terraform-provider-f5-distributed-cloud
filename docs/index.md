---
page_title: "F5 Distributed Cloud Provider"
subcategory: ""
description: |-
  Terraform provider for managing F5 Distributed Cloud resources.
---

# F5 Distributed Cloud Provider

The F5 Distributed Cloud (F5 XC) provider enables Terraform to manage F5 XC resources including namespaces, origin pools, HTTP load balancers, cloud sites, and application firewalls.

## Example Usage

```terraform
terraform {
  required_providers {
    f5_distributed_cloud = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 0.1"
    }
  }
}

# Configure the provider
provider "f5_distributed_cloud" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.api_token
}

# Create a namespace
resource "f5_distributed_cloud_namespace" "example" {
  name        = "my-namespace"
  description = "Example namespace"
}

# Create an origin pool
resource "f5_distributed_cloud_origin_pool" "example" {
  name              = "backend-pool"
  namespace         = f5_distributed_cloud_namespace.example.name
  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    public_name {
      dns_name = "backend.example.com"
    }
  }
}

# Create an HTTP load balancer
resource "f5_distributed_cloud_http_loadbalancer" "example" {
  name        = "frontend-lb"
  namespace   = f5_distributed_cloud_namespace.example.name
  domains     = ["app.example.com"]
  http_port   = 80

  advertise_on_public = true

  default_pool {
    name      = f5_distributed_cloud_origin_pool.example.name
    namespace = f5_distributed_cloud_namespace.example.name
  }
}
```

## Authentication

The provider supports two authentication methods:

### API Token (Recommended)

```terraform
provider "f5_distributed_cloud" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.api_token
}
```

Or via environment variables:

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
```

### P12 Certificate

```terraform
provider "f5_distributed_cloud" {
  api_url      = "https://your-tenant.console.ves.volterra.io/api"
  p12_file     = "/path/to/certificate.p12"
  p12_password = var.p12_password
}
```

Or via environment variables:

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_P12_FILE="/path/to/certificate.p12"
export F5XC_API_P12_PASSWORD="your-password"
```

## Schema

### Optional

- `api_url` (String) The URL for the F5 XC API. Can also be set via the `F5XC_API_URL` environment variable.
- `api_token` (String, Sensitive) API token for authentication. Can also be set via the `F5XC_API_TOKEN` environment variable.
- `p12_file` (String) Path to the P12 certificate file for authentication. Can also be set via the `F5XC_API_P12_FILE` environment variable.
- `p12_password` (String, Sensitive) Password for the P12 certificate file. Can also be set via the `F5XC_API_P12_PASSWORD` environment variable.
