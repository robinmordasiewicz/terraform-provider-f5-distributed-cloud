---
page_title: "f5xc_data_delivery Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Custom handler in Data Delivery microservice will forward request(s) to backend API(s)
---

# f5xc_data_delivery (Data Source)

Custom handler in Data Delivery microservice will forward request(s) to backend API(s)

## Example Usage

```hcl
data "f5xc_data_delivery" "example" {
  name      = "example-data_delivery"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
