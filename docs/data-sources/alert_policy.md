---
page_title: "f5_distributed_cloud_alert_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Alert Policy is used to specify a set of routes to match the incoming alert and the set of receivers to send the alert notification if there is a match. An Alert policy object defines a node in the...
---

# f5_distributed_cloud_alert_policy (Data Source)

Alert Policy is used to specify a set of routes to match the incoming alert and the set of receivers to send the alert notification if there is a match. An Alert policy object defines a node in the...

## Example Usage

```hcl
data "f5_distributed_cloud_alert_policy" "example" {
  name      = "example-alert_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
