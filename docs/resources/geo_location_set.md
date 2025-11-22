---
page_title: "f5xc_geo_location_set Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Defines the geo_location_set created by user
---

# f5xc_geo_location_set (Resource)

Defines the geo_location_set created by user

## Example Usage

```hcl
resource "f5xc_geo_location_set" "example" {
  name        = "example-geo_location_set"
  namespace   = "system"
  description = "Example GeoLocationSet resource"
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

GeoLocationSet can be imported using the namespace and name:

```shell
terraform import f5xc_geo_location_set.example namespace/name
```
