---
page_title: "f5_distributed_cloud_user_group Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Represents group for a given tenant
---

# f5_distributed_cloud_user_group (Data Source)

Represents group for a given tenant

## Example Usage

```hcl
data "f5_distributed_cloud_user_group" "example" {
  name      = "example-user_group"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
