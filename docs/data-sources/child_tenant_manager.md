---
page_title: "f5xc_child_tenant_manager Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Child Tenant Manager uses Tenant Profile template to create child tenant and store its object reference.  Any look up for child tenant should be done from child tenant manager since it stores objec...
---

# f5xc_child_tenant_manager (Data Source)

Child Tenant Manager uses Tenant Profile template to create child tenant and store its object reference.  Any look up for child tenant should be done from child tenant manager since it stores objec...

## Example Usage

```hcl
data "f5xc_child_tenant_manager" "example" {
  name      = "example-child_tenant_manager"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
