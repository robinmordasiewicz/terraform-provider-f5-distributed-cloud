---
page_title: "f5xc_geo_config Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Geo Config.
---

# f5xc_geo_config (Data Source)

Geo Config.

## Example Usage

```hcl
data "f5xc_geo_config" "example" {
  name      = "example-geo_config"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
