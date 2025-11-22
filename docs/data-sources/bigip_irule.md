---
page_title: "f5xc_bigip_irule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BIG-IP iRule Service manages iRule Library for customers
---

# f5xc_bigip_irule (Data Source)

BIG-IP iRule Service manages iRule Library for customers

## Example Usage

```hcl
data "f5xc_bigip_irule" "example" {
  name      = "example-bigip_irule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
