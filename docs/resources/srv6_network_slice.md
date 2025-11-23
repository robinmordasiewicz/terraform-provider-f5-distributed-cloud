---
page_title: "f5_distributed_cloud_srv6_network_slice Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  An srv6_network_slice represents a network slice in an operator network that uses SRv6.
---

# f5_distributed_cloud_srv6_network_slice (Resource)

An srv6_network_slice represents a network slice in an operator network that uses SRv6.

## Example Usage

```hcl
resource "f5_distributed_cloud_srv6_network_slice" "example" {
  name        = "example-srv6_network_slice"
  namespace   = "system"
  description = "Example Srv6NetworkSlice resource"
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

Srv6NetworkSlice can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_srv6_network_slice.example namespace/name
```
