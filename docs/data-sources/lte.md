---
page_title: "f5xc_lte Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Proto definitions for runtime LTE configuration on sites.
---

# f5xc_lte (Data Source)

Proto definitions for runtime LTE configuration on sites.

## Example Usage

```hcl
data "f5xc_lte" "example" {
  name      = "example-lte"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
