---
page_title: "f5_distributed_cloud_geo_location_set Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Defines the geo_location_set created by user
---

# f5_distributed_cloud_geo_location_set (Data Source)

Defines the geo_location_set created by user

## Example Usage

```hcl
data "f5_distributed_cloud_geo_location_set" "example" {
  name      = "example-geo_location_set"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
