---
page_title: "f5_distributed_cloud_debug Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Proto definitions for debugging site
---

# f5_distributed_cloud_debug (Data Source)

Proto definitions for debugging site

## Example Usage

```hcl
data "f5_distributed_cloud_debug" "example" {
  name      = "example-debug"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
