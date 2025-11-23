---
page_title: "f5_distributed_cloud_shape_bot_defense_instance Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Shape Bot Defense Instance is the main configuration for a Shape Integration. It defines various configuration parameters needed to use Shape SSEs.
---

# f5_distributed_cloud_shape_bot_defense_instance (Data Source)

Shape Bot Defense Instance is the main configuration for a Shape Integration. It defines various configuration parameters needed to use Shape SSEs.

## Example Usage

```hcl
data "f5_distributed_cloud_shape_bot_defense_instance" "example" {
  name      = "example-shape_bot_defense_instance"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
