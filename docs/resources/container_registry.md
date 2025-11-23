---
page_title: "f5_distributed_cloud_container_registry Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Container registry is the container or docker registry information
---

# f5_distributed_cloud_container_registry (Resource)

Container registry is the container or docker registry information

## Example Usage

```hcl
resource "f5_distributed_cloud_container_registry" "example" {
  name        = "example-container_registry"
  namespace   = "system"
  description = "Example ContainerRegistry resource"
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

ContainerRegistry can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_container_registry.example namespace/name
```
