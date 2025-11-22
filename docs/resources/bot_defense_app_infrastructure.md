---
page_title: "f5xc_bot_defense_app_infrastructure Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Bot Defense App Infrastructure is the main configuration for a Bot Defense Advanced Integration. It is intended to be created per tenant. It defines various configuration parameters needed to use B...
---

# f5xc_bot_defense_app_infrastructure (Resource)

Bot Defense App Infrastructure is the main configuration for a Bot Defense Advanced Integration. It is intended to be created per tenant. It defines various configuration parameters needed to use B...

## Example Usage

```hcl
resource "f5xc_bot_defense_app_infrastructure" "example" {
  name        = "example-bot_defense_app_infrastructure"
  namespace   = "system"
  description = "Example BotDefenseAppInfrastructure resource"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace where the resource will be created.
- `description` - (Optional) Description of the resource.
- `labels` - (Optional) Labels for the resource.
- `annotations` - (Optional) Annotations for the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.

## Import

BotDefenseAppInfrastructure can be imported using the namespace and name:

```shell
terraform import f5xc_bot_defense_app_infrastructure.example namespace/name
```
