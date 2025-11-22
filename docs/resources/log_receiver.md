---
page_title: "f5xc_log_receiver Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Log Receiver is used to specify a receiver (syslog, splunk, datadog etc.,) to send the log. Log receiver need to be reachable on site local network.
---

# f5xc_log_receiver (Resource)

Log Receiver is used to specify a receiver (syslog, splunk, datadog etc.,) to send the log. Log receiver need to be reachable on site local network.

## Example Usage

```hcl
resource "f5xc_log_receiver" "example" {
  name        = "example-log_receiver"
  namespace   = "system"
  description = "Example LogReceiver resource"
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

LogReceiver can be imported using the namespace and name:

```shell
terraform import f5xc_log_receiver.example namespace/name
```
