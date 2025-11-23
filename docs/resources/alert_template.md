---
page_title: "f5_distributed_cloud_alert_template Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BRM Alerts Alert Template
---

# f5_distributed_cloud_alert_template (Resource)

BRM Alerts Alert Template

## Example Usage

```hcl
resource "f5_distributed_cloud_alert_template" "example" {
  name        = "example-alert_template"
  namespace   = "system"
  description = "Example AlertTemplate resource"
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

AlertTemplate can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_alert_template.example namespace/name
```
