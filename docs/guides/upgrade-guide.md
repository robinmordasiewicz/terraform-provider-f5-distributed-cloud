# Upgrade Guide

This guide provides information for upgrading between versions of the F5 Distributed Cloud Terraform provider.

## Upgrading from 0.x to 1.0

### Breaking Changes

Version 1.0 introduces several breaking changes to align with Terraform best practices and F5 XC API updates.

#### Provider Configuration

No changes to provider configuration are required.

#### Resource Changes

##### f5xc_namespace

No breaking changes.

##### f5xc_origin_pool

- `endpoint_selection` has been replaced with `loadbalancer_algorithm`
- Default port behavior changed: port is now always required

**Migration:**

```hcl
# Before
resource "f5xc_origin_pool" "example" {
  name               = "example"
  namespace          = "default"
  endpoint_selection = "DISTRIBUTED"
  # port was optional with defaults
}

# After
resource "f5xc_origin_pool" "example" {
  name                   = "example"
  namespace              = "default"
  loadbalancer_algorithm = "ROUND_ROBIN"
  port                   = 80  # Now required
}
```

##### f5xc_http_loadbalancer

- `advertise_on_public_default_vip` renamed to `advertise_on_public`
- `default_route_pools` renamed to `default_pool`

**Migration:**

```hcl
# Before
resource "f5xc_http_loadbalancer" "example" {
  name                            = "example"
  namespace                       = "default"
  advertise_on_public_default_vip = true
  default_route_pools {
    pool {
      name = "pool"
    }
    weight = 1
  }
}

# After
resource "f5xc_http_loadbalancer" "example" {
  name                = "example"
  namespace           = "default"
  advertise_on_public = true
  default_pool {
    name = "pool"
  }
}
```

##### f5xc_cloud_site

- AWS, Azure, and GCP site resources have been unified into `f5xc_cloud_site`
- `site_type` and `cloud_provider` attributes are now required

**Migration:**

```hcl
# Before (AWS)
resource "volterra_aws_vpc_site" "example" {
  name      = "example"
  namespace = "system"
  aws_region = "us-east-1"
}

# After
resource "f5xc_cloud_site" "example" {
  name           = "example"
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "us-east-1"
}
```

##### f5xc_app_firewall

- `blocking_mode` renamed to `mode`
- Bot protection configuration structure changed

**Migration:**

```hcl
# Before
resource "f5xc_app_firewall" "example" {
  name          = "example"
  namespace     = "default"
  blocking_mode = true
}

# After
resource "f5xc_app_firewall" "example" {
  name      = "example"
  namespace = "default"
  mode      = "blocking"  # or "monitoring"
}
```

### Deprecations

The following features are deprecated and will be removed in version 2.0:

- None currently

### New Features

- Unified cloud site resource supporting AWS, Azure, and GCP
- Enhanced bot protection configuration
- Custom blocking page support for app firewall
- Improved error messages and validation

## Upgrade Process

### Step 1: Backup State

Always backup your state file before upgrading:

```bash
cp terraform.tfstate terraform.tfstate.backup
```

### Step 2: Update Provider Version

Update your provider version constraint:

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 1.0"
    }
  }
}
```

### Step 3: Update Configuration

Review your configuration against the breaking changes above and update as needed.

### Step 4: Initialize and Plan

```bash
terraform init -upgrade
terraform plan
```

Review the plan carefully for any unexpected changes.

### Step 5: Apply Changes

```bash
terraform apply
```

### Step 6: Verify

Verify your resources are functioning correctly after the upgrade.

## State Migration

In some cases, you may need to migrate state for renamed resources:

```bash
# Example: Migrate from old resource name to new
terraform state mv f5xc_aws_vpc_site.example f5xc_cloud_site.example
```

## Rollback

If you encounter issues:

1. Restore your state backup:
   ```bash
   cp terraform.tfstate.backup terraform.tfstate
   ```

2. Revert provider version:
   ```hcl
   terraform {
     required_providers {
       f5xc = {
         source  = "robinmordasiewicz/f5-distributed-cloud"
         version = "~> 0.x"  # Previous version
       }
     }
   }
   ```

3. Reinitialize:
   ```bash
   terraform init -upgrade
   ```

## Getting Help

If you encounter issues during the upgrade:

1. Check the [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/issues)
2. Review the [changelog](../../CHANGELOG.md)
3. Open a new issue if needed

## Version History

| Version | Release Date | Key Changes |
|---------|--------------|-------------|
| 0.1.0 | TBD | Initial release |
| 1.0.0 | TBD | First stable release |
