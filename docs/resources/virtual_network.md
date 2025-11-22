---
page_title: "f5xc_virtual_network Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual network is an isolated L3 network. A virtual network can contain one or more IP Subnets. Network elements in a virtual network (interface, host, listener etc...) can talk to each other dire...
---

# f5xc_virtual_network (Resource)

Virtual network is an isolated L3 network. A virtual network can contain one or more IP Subnets. Network elements in a virtual network (interface, host, listener etc...) can talk to each other dire...

## Example Usage

```hcl
resource "f5xc_virtual_network" "example" {
  name        = "example-virtual_network"
  namespace   = "system"
  description = "Example VirtualNetwork resource"
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

VirtualNetwork can be imported using the namespace and name:

```shell
terraform import f5xc_virtual_network.example namespace/name
```
