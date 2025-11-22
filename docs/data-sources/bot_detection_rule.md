---
page_title: "f5xc_bot_detection_rule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configures Bot Detection Rule
---

# f5xc_bot_detection_rule (Data Source)

Configures Bot Detection Rule

## Example Usage

```hcl
data "f5xc_bot_detection_rule" "example" {
  name      = "example-bot_detection_rule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
