---
page_title: "f5xc_navigation_tile Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Starting point of access to addon service from VoltConsole home page. Addon service is referenced to a navigation tile and a navigation tile can host one or more addon services.
---

# f5xc_navigation_tile (Data Source)

Starting point of access to addon service from VoltConsole home page. Addon service is referenced to a navigation tile and a navigation tile can host one or more addon services.

## Example Usage

```hcl
data "f5xc_navigation_tile" "example" {
  name      = "example-navigation_tile"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
