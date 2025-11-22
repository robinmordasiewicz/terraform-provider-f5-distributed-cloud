---
page_title: "f5xc_tenant Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Package for working with Tenant representation.
---

# f5xc_tenant (Data Source)

Package for working with Tenant representation.

## Example Usage

```hcl
data "f5xc_tenant" "example" {
  name      = "example-tenant"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
