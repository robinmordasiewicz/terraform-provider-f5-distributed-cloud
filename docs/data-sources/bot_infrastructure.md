---
page_title: "f5xc_bot_infrastructure Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configures Bot Infrastructure by bot infrastructure
---

# f5xc_bot_infrastructure (Data Source)

Configures Bot Infrastructure by bot infrastructure

## Example Usage

```hcl
data "f5xc_bot_infrastructure" "example" {
  name      = "example-bot_infrastructure"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
