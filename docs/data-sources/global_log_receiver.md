---
page_title: "f5_distributed_cloud_global_log_receiver Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Global Log Receiver is used to specify a receiver (s3 bucket, etc.) for periodic streaming of access logs
---

# f5_distributed_cloud_global_log_receiver (Data Source)

Global Log Receiver is used to specify a receiver (s3 bucket, etc.) for periodic streaming of access logs

## Example Usage

```hcl
data "f5_distributed_cloud_global_log_receiver" "example" {
  name      = "example-global_log_receiver"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
