---
page_title: "f5xc_srv6_network_slice Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  An srv6_network_slice represents a network slice in an operator network that uses SRv6.
---

# f5xc_srv6_network_slice (Data Source)

An srv6_network_slice represents a network slice in an operator network that uses SRv6.

## Example Usage

```hcl
data "f5xc_srv6_network_slice" "example" {
  name      = "example-srv6_network_slice"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
