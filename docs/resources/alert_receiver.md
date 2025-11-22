---
page_title: "f5xc_alert_receiver Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Alert Receiver is used to specify a receiver (slack, pagerDuty, etc.,) to send the alert notifications. An Alert Receiver may be associated with one or more Alert Policy objects, which defines one ...
---

# f5xc_alert_receiver (Resource)

Alert Receiver is used to specify a receiver (slack, pagerDuty, etc.,) to send the alert notifications. An Alert Receiver may be associated with one or more Alert Policy objects, which defines one ...

## Example Usage

```hcl
resource "f5xc_alert_receiver" "example" {
  name        = "example-alert_receiver"
  namespace   = "system"
  description = "Example AlertReceiver resource"
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

AlertReceiver can be imported using the namespace and name:

```shell
terraform import f5xc_alert_receiver.example namespace/name
```
