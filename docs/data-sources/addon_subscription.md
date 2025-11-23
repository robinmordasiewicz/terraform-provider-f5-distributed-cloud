---
page_title: "f5_distributed_cloud_addon_subscription Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Represents addon subscription  Eywa will create the schema.pbac.addon_subscription object (SUBSCRIPTION_PENDING) SRE/Support team member via f5xc-support tenant changes the status of the addon_subs...
---

# f5_distributed_cloud_addon_subscription (Data Source)

Represents addon subscription  Eywa will create the schema.pbac.addon_subscription object (SUBSCRIPTION_PENDING) SRE/Support team member via f5xc-support tenant changes the status of the addon_subs...

## Example Usage

```hcl
data "f5_distributed_cloud_addon_subscription" "example" {
  name      = "example-addon_subscription"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
