---
page_title: "f5_distributed_cloud_namespace Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  namespace creates logical independent workspace within a tenant. Within a namespace contained objects should have unique names.      Object within a namespace can only refer to objects within the s...
---

# f5_distributed_cloud_namespace (Data Source)

namespace creates logical independent workspace within a tenant. Within a namespace contained objects should have unique names.      Object within a namespace can only refer to objects within the s...

## Example Usage

```hcl
data "f5_distributed_cloud_namespace" "example" {
  name      = "example-namespace"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
