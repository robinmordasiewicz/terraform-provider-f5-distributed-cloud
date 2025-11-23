---
page_title: "f5_distributed_cloud_code_base_integration Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Code base integration
---

# f5_distributed_cloud_code_base_integration (Data Source)

Code base integration

## Example Usage

```hcl
data "f5_distributed_cloud_code_base_integration" "example" {
  name      = "example-code_base_integration"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
