---
page_title: "f5_distributed_cloud_invoice Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Invoice listing and downloading APIs
---

# f5_distributed_cloud_invoice (Data Source)

Invoice listing and downloading APIs

## Example Usage

```hcl
data "f5_distributed_cloud_invoice" "example" {
  name      = "example-invoice"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
