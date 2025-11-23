---
page_title: "f5_distributed_cloud_mobile_base_config Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configures Mobile SDK Base Configuration
---

# f5_distributed_cloud_mobile_base_config (Data Source)

Configures Mobile SDK Base Configuration

## Example Usage

```hcl
data "f5_distributed_cloud_mobile_base_config" "example" {
  name      = "example-mobile_base_config"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
