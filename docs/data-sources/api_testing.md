---
page_title: "f5xc_api_testing Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  This is the api testing type
---

# f5xc_api_testing (Data Source)

This is the api testing type

## Example Usage

```hcl
data "f5xc_api_testing" "example" {
  name      = "example-api_testing"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
