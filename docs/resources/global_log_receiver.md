---
page_title: "f5xc_global_log_receiver Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Global Log Receiver is used to specify a receiver (s3 bucket, etc.) for periodic streaming of access logs
---

# f5xc_global_log_receiver (Resource)

Global Log Receiver is used to specify a receiver (s3 bucket, etc.) for periodic streaming of access logs

## Example Usage

```hcl
resource "f5xc_global_log_receiver" "example" {
  name        = "example-global_log_receiver"
  namespace   = "system"
  description = "Example GlobalLogReceiver resource"
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

GlobalLogReceiver can be imported using the namespace and name:

```shell
terraform import f5xc_global_log_receiver.example namespace/name
```
