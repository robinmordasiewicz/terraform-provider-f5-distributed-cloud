---
page_title: "f5xc_allowed_tenant Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Allowed tenant object will allow tenant in the name field to manage tenant in which its created. Admin can create allowed_tenant configuration if the tenant needs to be managed by tenant in allowed...
---

# f5xc_allowed_tenant (Resource)

Allowed tenant object will allow tenant in the name field to manage tenant in which its created. Admin can create allowed_tenant configuration if the tenant needs to be managed by tenant in allowed...

## Example Usage

```hcl
resource "f5xc_allowed_tenant" "example" {
  name        = "example-allowed_tenant"
  namespace   = "system"
  description = "Example AllowedTenant resource"
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

AllowedTenant can be imported using the namespace and name:

```shell
terraform import f5xc_allowed_tenant.example namespace/name
```
