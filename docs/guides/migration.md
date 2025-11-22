# Migration Guide

This guide helps you migrate from the official F5 XC Terraform provider to this community provider.

## Overview

This community provider is designed to be compatible with the official F5 XC provider while offering enhanced features and simplified workflows. Most resources can be migrated with minimal configuration changes.

## Prerequisites

- Terraform 1.0 or later
- Go 1.21 or later (if building from source)
- F5 XC API credentials (API token or P12 certificate)

## Migration Steps

### Step 1: Update Provider Configuration

Change your provider block from:

```hcl
terraform {
  required_providers {
    volterra = {
      source  = "volterraedge/volterra"
      version = "~> 0.11"
    }
  }
}

provider "volterra" {
  api_p12_file = var.p12_file
  url          = var.api_url
}
```

To:

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 0.1"
    }
  }
}

provider "f5xc" {
  api_url   = var.api_url
  api_token = var.api_token  # or use p12_file/p12_password
}
```

### Step 2: Update Resource Names

Update resource type names from `volterra_*` to `f5xc_*`:

| Old Resource Name | New Resource Name |
|-------------------|-------------------|
| `volterra_namespace` | `f5xc_namespace` |
| `volterra_origin_pool` | `f5xc_origin_pool` |
| `volterra_http_loadbalancer` | `f5xc_http_loadbalancer` |
| `volterra_app_firewall` | `f5xc_app_firewall` |
| `volterra_aws_vpc_site` | `f5xc_cloud_site` |
| `volterra_azure_vnet_site` | `f5xc_cloud_site` |
| `volterra_gcp_vpc_site` | `f5xc_cloud_site` |

### Step 3: Update Resource Attributes

Most attributes remain the same, but some have been renamed for clarity:

#### Origin Pool

```hcl
# Old
resource "volterra_origin_pool" "example" {
  name                   = "example"
  namespace              = "default"
  endpoint_selection     = "DISTRIBUTED"
  origin_servers {
    public_ip {
      ip = "192.168.1.1"
    }
  }
  port = 80
}

# New
resource "f5xc_origin_pool" "example" {
  name        = "example"
  namespace   = "default"
  port        = 80
  endpoint_protocol = "http"
  origin_servers {
    public_ip {
      ip = "192.168.1.1"
    }
  }
}
```

#### HTTP Load Balancer

```hcl
# Old
resource "volterra_http_loadbalancer" "example" {
  name                            = "example"
  namespace                       = "default"
  domains                         = ["example.com"]
  advertise_on_public_default_vip = true
  default_route_pools {
    pool {
      name      = volterra_origin_pool.example.name
      namespace = "default"
    }
    weight = 1
  }
}

# New
resource "f5xc_http_loadbalancer" "example" {
  name        = "example"
  namespace   = "default"
  domains     = ["example.com"]
  http_port   = 80
  advertise_on_public = true
  default_pool {
    name      = f5xc_origin_pool.example.name
    namespace = "default"
  }
}
```

#### Cloud Sites

Cloud sites have been unified into a single resource type:

```hcl
# Old (AWS)
resource "volterra_aws_vpc_site" "example" {
  name      = "aws-site"
  namespace = "system"
  # ...
}

# New
resource "f5xc_cloud_site" "example" {
  name           = "aws-site"
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "us-east-1"
  # ...
}
```

### Step 4: Import Existing Resources

After updating your configuration, import existing resources:

```bash
# Import namespace
terraform import f5xc_namespace.example my-namespace

# Import origin pool
terraform import f5xc_origin_pool.example my-namespace/my-origin-pool

# Import HTTP load balancer
terraform import f5xc_http_loadbalancer.example my-namespace/my-lb

# Import app firewall
terraform import f5xc_app_firewall.example my-namespace/my-firewall

# Import cloud site
terraform import f5xc_cloud_site.example my-site
```

### Step 5: Verify State

Run `terraform plan` to verify the migration was successful:

```bash
terraform plan
```

The plan should show no changes if the migration was performed correctly.

## Authentication Migration

### From P12 Certificate to API Token

If you're migrating from certificate-based authentication to API token:

1. Generate an API token in the F5 XC console
2. Update your provider configuration:

```hcl
provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.f5xc_api_token
}
```

3. Set the environment variable:

```bash
export F5XC_API_TOKEN="your-api-token"
```

### Continuing with P12 Certificate

If you prefer to continue using P12 certificates:

```hcl
provider "f5xc" {
  api_url      = "https://your-tenant.console.ves.volterra.io/api"
  p12_file     = var.p12_file
  p12_password = var.p12_password
}
```

## Troubleshooting

### Common Issues

#### Resource Not Found During Import

Ensure the resource exists in F5 XC and you're using the correct import ID format:
- Namespaced resources: `namespace/name`
- System resources (like cloud sites): `name`

#### Authentication Errors

Verify your credentials are correct and have the necessary permissions:
- API tokens must have appropriate role assignments
- P12 certificates must not be expired

#### State Drift

If you see unexpected changes in `terraform plan`:
1. Check if resource attributes have been renamed
2. Verify default values match your expected configuration
3. Review the resource documentation for any changes

### Getting Help

- Check the [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/issues)
- Review the [examples](../examples/) directory
- Consult the [F5 XC documentation](https://docs.cloud.f5.com/)

## Version Compatibility

| Provider Version | Terraform Version | F5 XC API Version |
|------------------|-------------------|-------------------|
| 0.1.x | >= 1.0 | v1 |

## Rollback Plan

If you need to rollback to the original provider:

1. Keep a backup of your state file before migration
2. Revert provider configuration changes
3. Run `terraform init -upgrade` to reinstall the original provider
4. Restore the backup state file if needed
