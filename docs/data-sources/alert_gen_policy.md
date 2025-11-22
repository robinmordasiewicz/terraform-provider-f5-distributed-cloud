---
page_title: "f5xc_alert_gen_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BRM Alerts Alert Generation Policy
---

# f5xc_alert_gen_policy (Data Source)

BRM Alerts Alert Generation Policy

## Example Usage

```hcl
data "f5xc_alert_gen_policy" "example" {
  name      = "example-alert_gen_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
