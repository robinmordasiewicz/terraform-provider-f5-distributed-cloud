---
page_title: "f5xc_alert_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Alert Policy is used to specify a set of routes to match the incoming alert and the set of receivers to send the alert notification if there is a match. An Alert policy object defines a node in the...
---

# f5xc_alert_policy (Resource)

Alert Policy is used to specify a set of routes to match the incoming alert and the set of receivers to send the alert notification if there is a match. An Alert policy object defines a node in the...

## Example Usage

```hcl
resource "f5xc_alert_policy" "example" {
  name        = "example-alert_policy"
  namespace   = "system"
  description = "Example AlertPolicy resource"
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

AlertPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_alert_policy.example namespace/name
```
