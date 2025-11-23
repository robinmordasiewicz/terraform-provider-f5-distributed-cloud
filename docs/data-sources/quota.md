---
page_title: "f5_distributed_cloud_quota Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Quota object is used to configure the limits on how many of a resource type can be in use by a tenant
---

# f5_distributed_cloud_quota (Data Source)

Quota object is used to configure the limits on how many of a resource type can be in use by a tenant

## Example Usage

```hcl
data "f5_distributed_cloud_quota" "example" {
  name      = "example-quota"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
