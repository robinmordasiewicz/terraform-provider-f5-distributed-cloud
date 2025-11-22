# Migration Guide: Volterra Provider to F5 Distributed Cloud Provider

This guide demonstrates how to migrate from the `volterraedge/volterra` provider to the
`robinmordasiewicz/f5-distributed-cloud` provider.

## Overview

The F5 Distributed Cloud provider is a modern, auto-generated provider that maintains
compatibility with the Volterra provider while providing additional features:

- **175 resources** vs 157 in volterra provider
- **229 data sources** vs 6 in volterra provider
- Auto-generated from official F5 XC OpenAPI specifications
- Full coverage of all core F5 XC resources

## Key Differences

| Aspect | Volterra Provider | F5 Distributed Cloud Provider |
|--------|-------------------|------------------------------|
| Resource prefix | `volterra_` | `f5xc_` |
| Provider name | `volterraedge/volterra` | `robinmordasiewicz/f5-distributed-cloud` |
| Authentication | Same | Same (token or certificate) |
| API URL format | Same | Same |

## Migration Steps

### Step 1: Update Provider Block

**Before (Volterra):**
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
  api_p12_file = var.api_p12_file
  url          = var.api_url
}
```

**After (F5 Distributed Cloud):**
```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 1.0"
    }
  }
}

provider "f5xc" {
  api_url  = var.api_url
  p12_file = var.api_p12_file
  # Or use token authentication:
  # api_token = var.api_token
}
```

### Step 2: Update Resource References

Resource names change from `volterra_*` to `f5xc_*`:

| Volterra Resource | F5 XC Resource |
|-------------------|----------------|
| `volterra_namespace` | `f5xc_namespace` |
| `volterra_origin_pool` | `f5xc_origin_pool` |
| `volterra_http_loadbalancer` | `f5xc_http_loadbalancer` |
| `volterra_app_firewall` | `f5xc_app_firewall` |
| `volterra_healthcheck` | `f5xc_healthcheck` |
| `volterra_virtual_site` | `f5xc_virtual_site` |
| `volterra_aws_vpc_site` | `f5xc_aws_vpc_site` |
| `volterra_azure_vnet_site` | `f5xc_azure_vnet_site` |
| `volterra_gcp_vpc_site` | `f5xc_gcp_vpc_site` |

### Step 3: State Migration

To migrate existing state without recreating resources:

1. **Export current state:**
   ```bash
   terraform state pull > terraform.tfstate.backup
   ```

2. **Update provider and resources in your .tf files**

3. **Use state mv to rename resources:**
   ```bash
   # Example: Rename namespace resource
   terraform state mv volterra_namespace.my_ns f5xc_namespace.my_ns

   # Example: Rename origin pool
   terraform state mv volterra_origin_pool.my_pool f5xc_origin_pool.my_pool
   ```

4. **Re-initialize with new provider:**
   ```bash
   terraform init -upgrade
   ```

5. **Verify with plan:**
   ```bash
   terraform plan
   # Should show no changes if migration was successful
   ```

## Example Migration

See the `before/` and `after/` directories for complete working examples.
