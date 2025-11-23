---
page_title: "f5_distributed_cloud_quota Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Quota object is used to configure the limits on how many of a resource type can be in use by a tenant
---

# f5_distributed_cloud_quota (Resource)

Quota object is used to configure the limits on how many of a resource type can be in use by a tenant

## Example Usage

```hcl
resource "f5_distributed_cloud_quota" "example" {
  name        = "example-quota"
  namespace   = "system"
  description = "Example Quota resource"
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

Quota can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_quota.example namespace/name
```
