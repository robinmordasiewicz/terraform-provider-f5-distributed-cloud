---
page_title: "f5xc_subscription Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Subscription details APIs.
---

# f5xc_subscription (Data Source)

Subscription details APIs.

## Example Usage

```hcl
data "f5xc_subscription" "example" {
  name      = "example-subscription"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
