---
page_title: "f5xc_irule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  iRule object defines the iRule that can be used in CRUD, to create and manage iRule for manipulating the application traffic.
---

# f5xc_irule (Data Source)

iRule object defines the iRule that can be used in CRUD, to create and manage iRule for manipulating the application traffic.

## Example Usage

```hcl
data "f5xc_irule" "example" {
  name      = "example-irule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
