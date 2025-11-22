---
page_title: "f5xc_upgrade_status Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Upgrade status custom APIs
---

# f5xc_upgrade_status (Data Source)

Upgrade status custom APIs

## Example Usage

```hcl
data "f5xc_upgrade_status" "example" {
  name      = "example-upgrade_status"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
