---
page_title: "f5_distributed_cloud_bot_endpoint_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configures Bot Endpoint Policy
---

# f5_distributed_cloud_bot_endpoint_policy (Data Source)

Configures Bot Endpoint Policy

## Example Usage

```hcl
data "f5_distributed_cloud_bot_endpoint_policy" "example" {
  name      = "example-bot_endpoint_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
