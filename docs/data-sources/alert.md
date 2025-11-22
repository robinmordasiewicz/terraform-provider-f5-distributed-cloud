---
page_title: "f5xc_alert Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Alert may be generated based on the metrics or based on severity level in the logs. All alerts are scoped by  tenant and namespace and tagged with the following default labels that can be used to f...
---

# f5xc_alert (Data Source)

Alert may be generated based on the metrics or based on severity level in the logs. All alerts are scoped by  tenant and namespace and tagged with the following default labels that can be used to f...

## Example Usage

```hcl
data "f5xc_alert" "example" {
  name      = "example-alert"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
