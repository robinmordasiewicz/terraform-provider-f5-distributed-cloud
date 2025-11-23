---
page_title: "f5_distributed_cloud_mobile_base_config Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configures Mobile SDK Base Configuration
---

# f5_distributed_cloud_mobile_base_config (Resource)

Configures Mobile SDK Base Configuration

## Example Usage

```hcl
resource "f5_distributed_cloud_mobile_base_config" "example" {
  name        = "example-mobile_base_config"
  namespace   = "system"
  description = "Example MobileBaseConfig resource"
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

MobileBaseConfig can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_mobile_base_config.example namespace/name
```
