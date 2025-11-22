---
page_title: "f5xc_tenant Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Package for working with Tenant representation.
---

# f5xc_tenant (Resource)

Package for working with Tenant representation.

## Example Usage

```hcl
resource "f5xc_tenant" "example" {
  name        = "example-tenant"
  namespace   = "system"
  description = "Example Tenant resource"
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

Tenant can be imported using the namespace and name:

```shell
terraform import f5xc_tenant.example namespace/name
```
