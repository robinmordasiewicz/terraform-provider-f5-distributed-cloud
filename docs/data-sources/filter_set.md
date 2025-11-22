---
page_title: "f5xc_filter_set Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Filter Set is a set of saved filtering criteria used in the Console. This allows users to declare named sets of filters so that they can be consistently used and shared to quickly reactivate a part...
---

# f5xc_filter_set (Data Source)

Filter Set is a set of saved filtering criteria used in the Console. This allows users to declare named sets of filters so that they can be consistently used and shared to quickly reactivate a part...

## Example Usage

```hcl
data "f5xc_filter_set" "example" {
  name      = "example-filter_set"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
