---
page_title: "f5_distributed_cloud_infraprotect_information Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Infraprotect information about the current organisation
---

# f5_distributed_cloud_infraprotect_information (Data Source)

Infraprotect information about the current organisation

## Example Usage

```hcl
data "f5_distributed_cloud_infraprotect_information" "example" {
  name      = "example-infraprotect_information"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
