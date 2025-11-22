---
page_title: "f5xc_static_component Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  stores information about the UI Components in key-value pair
---

# f5xc_static_component (Data Source)

stores information about the UI Components in key-value pair

## Example Usage

```hcl
data "f5xc_static_component" "example" {
  name      = "example-static_component"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
