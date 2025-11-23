---
page_title: "f5_distributed_cloud_child_tenant Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Child Tenant
---

# f5_distributed_cloud_child_tenant (Data Source)

Child Tenant

## Example Usage

```hcl
data "f5_distributed_cloud_child_tenant" "example" {
  name      = "example-child_tenant"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
