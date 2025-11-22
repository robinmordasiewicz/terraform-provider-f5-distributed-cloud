---
page_title: "f5xc_managed_tenant Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Managed tenant objects are required for declaring intent to manage a tenant. The tenant which is being managed is called a `Managed Tenant` or `MT`` and the tenant which is initiating the managemen...
---

# f5xc_managed_tenant (Resource)

Managed tenant objects are required for declaring intent to manage a tenant. The tenant which is being managed is called a `Managed Tenant` or `MT`` and the tenant which is initiating the managemen...

## Example Usage

```hcl
resource "f5xc_managed_tenant" "example" {
  name        = "example-managed_tenant"
  namespace   = "system"
  description = "Example ManagedTenant resource"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace where the resource will be created.
- `description` - (Optional) Description of the resource.
- `labels` - (Optional) Labels for the resource.
- `annotations` - (Optional) Annotations for the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.

## Import

ManagedTenant can be imported using the namespace and name:

```shell
terraform import f5xc_managed_tenant.example namespace/name
```
