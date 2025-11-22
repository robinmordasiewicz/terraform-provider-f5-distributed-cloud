---
page_title: "f5xc_setting Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Custom API of user settings.
---

# f5xc_setting (Data Source)

Custom API of user settings.

## Example Usage

```hcl
data "f5xc_setting" "example" {
  name      = "example-setting"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
