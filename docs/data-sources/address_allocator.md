---
page_title: "f5xc_address_allocator Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Address Allocator object is used to allocate an address or a subnet from a given address pool. Mode of the allocator object determines if the allocation is limited to VERs within the local site / c...
---

# f5xc_address_allocator (Data Source)

Address Allocator object is used to allocate an address or a subnet from a given address pool. Mode of the allocator object determines if the allocation is limited to VERs within the local site / c...

## Example Usage

```hcl
data "f5xc_address_allocator" "example" {
  name      = "example-address_allocator"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
