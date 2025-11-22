---
page_title: "f5xc_alert_gen_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BRM Alerts Alert Generation Policy
---

# f5xc_alert_gen_policy (Resource)

BRM Alerts Alert Generation Policy

## Example Usage

```hcl
resource "f5xc_alert_gen_policy" "example" {
  name        = "example-alert_gen_policy"
  namespace   = "system"
  description = "Example AlertGenPolicy resource"
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

AlertGenPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_alert_gen_policy.example namespace/name
```
