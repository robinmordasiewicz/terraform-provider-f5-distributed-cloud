---
page_title: "f5xc_receiver Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Data Delivery is used to specify a receiver (s3 bucket, etc.) for periodic delivery of customer/tenant data to customer sinks(destinations)
---

# f5xc_receiver (Data Source)

Data Delivery is used to specify a receiver (s3 bucket, etc.) for periodic delivery of customer/tenant data to customer sinks(destinations)

## Example Usage

```hcl
data "f5xc_receiver" "example" {
  name      = "example-receiver"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
