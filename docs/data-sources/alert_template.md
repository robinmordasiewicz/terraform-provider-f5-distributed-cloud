---
page_title: "f5xc_alert_template Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BRM Alerts Alert Template
---

# f5xc_alert_template (Data Source)

BRM Alerts Alert Template

## Example Usage

```hcl
data "f5xc_alert_template" "example" {
  name      = "example-alert_template"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
