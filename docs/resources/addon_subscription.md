---
page_title: "f5xc_addon_subscription Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Represents addon subscription  Eywa will create the schema.pbac.addon_subscription object (SUBSCRIPTION_PENDING) SRE/Support team member via f5xc-support tenant changes the status of the addon_subs...
---

# f5xc_addon_subscription (Resource)

Represents addon subscription  Eywa will create the schema.pbac.addon_subscription object (SUBSCRIPTION_PENDING) SRE/Support team member via f5xc-support tenant changes the status of the addon_subs...

## Example Usage

```hcl
resource "f5xc_addon_subscription" "example" {
  name        = "example-addon_subscription"
  namespace   = "system"
  description = "Example AddonSubscription resource"
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

AddonSubscription can be imported using the namespace and name:

```shell
terraform import f5xc_addon_subscription.example namespace/name
```
