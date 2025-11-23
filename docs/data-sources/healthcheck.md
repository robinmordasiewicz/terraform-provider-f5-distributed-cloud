---
page_title: "f5_distributed_cloud_healthcheck Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Health check configuration for a given cluster.
---

# f5_distributed_cloud_healthcheck (Data Source)

Health check configuration for a given cluster.

## Example Usage

```hcl
data "f5_distributed_cloud_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
