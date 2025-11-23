---
page_title: "f5_distributed_cloud_container_registry Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Container registry is the container or docker registry information
---

# f5_distributed_cloud_container_registry (Data Source)

Container registry is the container or docker registry information

## Example Usage

```hcl
data "f5_distributed_cloud_container_registry" "example" {
  name      = "example-container_registry"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
