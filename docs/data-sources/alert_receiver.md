---
page_title: "f5_distributed_cloud_alert_receiver Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Alert Receiver is used to specify a receiver (slack, pagerDuty, etc.,) to send the alert notifications. An Alert Receiver may be associated with one or more Alert Policy objects, which defines one ...
---

# f5_distributed_cloud_alert_receiver (Data Source)

Alert Receiver is used to specify a receiver (slack, pagerDuty, etc.,) to send the alert notifications. An Alert Receiver may be associated with one or more Alert Policy objects, which defines one ...

## Example Usage

```hcl
data "f5_distributed_cloud_alert_receiver" "example" {
  name      = "example-alert_receiver"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
