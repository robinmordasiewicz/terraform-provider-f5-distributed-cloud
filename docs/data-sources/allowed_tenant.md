---
page_title: "f5xc_allowed_tenant Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Allowed tenant object will allow tenant in the name field to manage tenant in which its created. Admin can create allowed_tenant configuration if the tenant needs to be managed by tenant in allowed...
---

# f5xc_allowed_tenant (Data Source)

Allowed tenant object will allow tenant in the name field to manage tenant in which its created. Admin can create allowed_tenant configuration if the tenant needs to be managed by tenant in allowed...

## Example Usage

```hcl
data "f5xc_allowed_tenant" "example" {
  name      = "example-allowed_tenant"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
