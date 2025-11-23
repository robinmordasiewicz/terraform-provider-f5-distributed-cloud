---
page_title: "f5_distributed_cloud_log_receiver Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Log Receiver is used to specify a receiver (syslog, splunk, datadog etc.,) to send the log. Log receiver need to be reachable on site local network.
---

# f5_distributed_cloud_log_receiver (Data Source)

Log Receiver is used to specify a receiver (syslog, splunk, datadog etc.,) to send the log. Log receiver need to be reachable on site local network.

## Example Usage

```hcl
data "f5_distributed_cloud_log_receiver" "example" {
  name      = "example-log_receiver"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
