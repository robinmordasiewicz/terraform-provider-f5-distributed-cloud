---
page_title: "f5_distributed_cloud_address_allocator Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Address Allocator object is used to allocate an address or a subnet from a given address pool. Mode of the allocator object determines if the allocation is limited to VERs within the local site / c...
---

# f5_distributed_cloud_address_allocator (Resource)

Address Allocator object is used to allocate an address or a subnet from a given address pool. Mode of the allocator object determines if the allocation is limited to VERs within the local site / c...

## Example Usage

```hcl
resource "f5_distributed_cloud_address_allocator" "example" {
  name        = "example-address_allocator"
  namespace   = "system"
  description = "Example AddressAllocator resource"
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

AddressAllocator can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_address_allocator.example namespace/name
```
