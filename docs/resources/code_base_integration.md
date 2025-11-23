---
page_title: "f5_distributed_cloud_code_base_integration Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Code base integration
---

# f5_distributed_cloud_code_base_integration (Resource)

Code base integration

## Example Usage

```hcl
resource "f5_distributed_cloud_code_base_integration" "example" {
  name        = "example-code_base_integration"
  namespace   = "system"
  description = "Example CodeBaseIntegration resource"
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

CodeBaseIntegration can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_code_base_integration.example namespace/name
```
