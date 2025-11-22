---
page_title: "f5xc_user Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  This API can be used to manage various attributes of the user like role, contact information etc.
---

# f5xc_user (Data Source)

This API can be used to manage various attributes of the user like role, contact information etc.

## Example Usage

```hcl
data "f5xc_user" "example" {
  name      = "example-user"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
