---
page_title: "f5xc_managed_tenant Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Managed tenant objects are required for declaring intent to manage a tenant. The tenant which is being managed is called a `Managed Tenant` or `MT`` and the tenant which is initiating the managemen...
---

# f5xc_managed_tenant (Data Source)

Managed tenant objects are required for declaring intent to manage a tenant. The tenant which is being managed is called a `Managed Tenant` or `MT`` and the tenant which is initiating the managemen...

## Example Usage

```hcl
data "f5xc_managed_tenant" "example" {
  name      = "example-managed_tenant"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
