---
page_title: "f5xc_bot_defense_app_infrastructure Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Bot Defense App Infrastructure is the main configuration for a Bot Defense Advanced Integration. It is intended to be created per tenant. It defines various configuration parameters needed to use B...
---

# f5xc_bot_defense_app_infrastructure (Data Source)

Bot Defense App Infrastructure is the main configuration for a Bot Defense Advanced Integration. It is intended to be created per tenant. It defines various configuration parameters needed to use B...

## Example Usage

```hcl
data "f5xc_bot_defense_app_infrastructure" "example" {
  name      = "example-bot_defense_app_infrastructure"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
