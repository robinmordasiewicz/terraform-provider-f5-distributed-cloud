---
page_title: "f5xc_receiver Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Data Delivery is used to specify a receiver (s3 bucket, etc.) for periodic delivery of customer/tenant data to customer sinks(destinations)
---

# f5xc_receiver (Resource)

Data Delivery is used to specify a receiver (s3 bucket, etc.) for periodic delivery of customer/tenant data to customer sinks(destinations)

## Example Usage

```hcl
resource "f5xc_receiver" "example" {
  name        = "example-receiver"
  namespace   = "system"
  description = "Example Receiver resource"
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

Receiver can be imported using the namespace and name:

```shell
terraform import f5xc_receiver.example namespace/name
```
