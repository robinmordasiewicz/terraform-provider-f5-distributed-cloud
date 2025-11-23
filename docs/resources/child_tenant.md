---
page_title: "f5_distributed_cloud_child_tenant Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Child Tenant
---

# f5_distributed_cloud_child_tenant (Resource)

Child Tenant

## Example Usage

```hcl
resource "f5_distributed_cloud_child_tenant" "example" {
  name        = "example-child_tenant"
  namespace   = "system"
  description = "Example ChildTenant resource"
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

ChildTenant can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_child_tenant.example namespace/name
```
