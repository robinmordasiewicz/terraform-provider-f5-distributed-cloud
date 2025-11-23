---
page_title: "f5_distributed_cloud_virtual_network Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual network is an isolated L3 network. A virtual network can contain one or more IP Subnets. Network elements in a virtual network (interface, host, listener etc...) can talk to each other dire...
---

# f5_distributed_cloud_virtual_network (Data Source)

Virtual network is an isolated L3 network. A virtual network can contain one or more IP Subnets. Network elements in a virtual network (interface, host, listener etc...) can talk to each other dire...

## Example Usage

```hcl
data "f5_distributed_cloud_virtual_network" "example" {
  name      = "example-virtual_network"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
