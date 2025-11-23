# Getting Started with F5 Distributed Cloud Terraform Provider

This guide walks you through setting up and using the F5 Distributed Cloud (F5 XC) Terraform provider.

## Prerequisites

- Terraform >= 1.0
- F5 Distributed Cloud tenant with API access
- API credentials (API token or P12 certificate)

## Installation

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    f5_distributed_cloud = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 0.1"
    }
  }
}
```

Run `terraform init` to download the provider.

## Authentication

The provider supports two authentication methods:

### API Token (Recommended)

```hcl
provider "f5_distributed_cloud" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.f5_distributed_cloud_api_token
}
```

Set the API token as an environment variable:

```bash
export F5XC_API_TOKEN="your-api-token"
```

### P12 Certificate

```hcl
provider "f5_distributed_cloud" {
  api_url      = "https://your-tenant.console.ves.volterra.io/api"
  p12_file     = "/path/to/certificate.p12"
  p12_password = var.p12_password
}
```

Set credentials via environment variables:

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_P12_FILE="/path/to/certificate.p12"
export F5XC_API_P12_PASSWORD="your-password"
```

## Basic Usage Example

Here's a complete example that creates a namespace, origin pool, and HTTP load balancer:

```hcl
terraform {
  required_providers {
    f5_distributed_cloud = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 0.1"
    }
  }
}

provider "f5_distributed_cloud" {
  api_url   = var.api_url
  api_token = var.api_token
}

# Create a namespace
resource "f5_distributed_cloud_namespace" "example" {
  name        = "my-app-namespace"
  description = "Namespace for my application"
}

# Create an origin pool
resource "f5_distributed_cloud_origin_pool" "backend" {
  name              = "backend-pool"
  namespace         = f5_distributed_cloud_namespace.example.name
  description       = "Backend server pool"
  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    public_name {
      dns_name = "backend.example.com"
    }
  }
}

# Create an HTTP load balancer
resource "f5_distributed_cloud_http_loadbalancer" "frontend" {
  name        = "frontend-lb"
  namespace   = f5_distributed_cloud_namespace.example.name
  description = "Frontend load balancer"
  domains     = ["app.example.com"]
  http_port   = 80

  advertise_on_public = true

  default_pool {
    name      = f5_distributed_cloud_origin_pool.backend.name
    namespace = f5_distributed_cloud_namespace.example.name
  }
}

# Add application firewall
resource "f5_distributed_cloud_app_firewall" "waf" {
  name        = "app-waf"
  namespace   = f5_distributed_cloud_namespace.example.name
  description = "Application firewall policy"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    violation_rating               = "medium"
  }

  bot_protection {
    good_bot_action       = "allow"
    malicious_bot_action  = "block"
    suspicious_bot_action = "challenge"
  }
}
```

## Variables

Create a `variables.tf` file:

```hcl
variable "api_url" {
  description = "F5 XC API URL"
  type        = string
}

variable "api_token" {
  description = "F5 XC API token"
  type        = string
  sensitive   = true
}
```

And a `terraform.tfvars` file (don't commit this!):

```hcl
api_url   = "https://your-tenant.console.ves.volterra.io/api"
api_token = "your-api-token"
```

## Applying the Configuration

```bash
# Initialize Terraform
terraform init

# Preview changes
terraform plan

# Apply changes
terraform apply

# Destroy resources when done
terraform destroy
```

## Best Practices

### Use Environment Variables for Credentials

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
```

Then simplify your provider configuration:

```hcl
provider "f5_distributed_cloud" {}
```

### Use Variables for Reusability

```hcl
variable "environment" {
  description = "Environment name"
  type        = string
  default     = "dev"
}

resource "f5_distributed_cloud_namespace" "env" {
  name        = "${var.environment}-namespace"
  description = "Namespace for ${var.environment} environment"
}
```

### Organize Resources with Modules

```
modules/
  load-balancer/
    main.tf
    variables.tf
    outputs.tf
  origin-pool/
    main.tf
    variables.tf
    outputs.tf
environments/
  dev/
    main.tf
  prod/
    main.tf
```

### Use Data Sources for Existing Resources

```hcl
# Reference existing namespace
data "f5_distributed_cloud_namespace" "existing" {
  name = "existing-namespace"
}

resource "f5_distributed_cloud_origin_pool" "pool" {
  name      = "my-pool"
  namespace = data.f5_distributed_cloud_namespace.existing.name
  # ...
}
```

## Next Steps

- Explore the [resource documentation](../resources/)
- Check out the [examples](../examples/)
- Read the [migration guide](./migration.md) if migrating from another provider
- Learn about [cloud sites](../resources/f5_distributed_cloud_cloud_site.md) for deploying infrastructure
