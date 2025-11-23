---
page_title: "f5_distributed_cloud_lma_region Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  LMA Region.
---

# f5_distributed_cloud_lma_region (Data Source)

LMA Region.

## Example Usage

```hcl
data "f5_distributed_cloud_lma_region" "example" {
  name      = "example-lma_region"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
