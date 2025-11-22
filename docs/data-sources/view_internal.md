---
page_title: "f5xc_view_internal Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  View Internal object contains reference to child objects of a given view
---

# f5xc_view_internal (Data Source)

View Internal object contains reference to child objects of a given view

## Example Usage

```hcl
data "f5xc_view_internal" "example" {
  name      = "example-view_internal"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
