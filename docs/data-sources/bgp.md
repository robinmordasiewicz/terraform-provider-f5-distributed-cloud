---
page_title: "f5_distributed_cloud_bgp Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BGP object represents configuration of bgp protocol on given network interface on customer edge site. It is template configuration that can be applied on a set of sites represented by list of sites...
---

# f5_distributed_cloud_bgp (Data Source)

BGP object represents configuration of bgp protocol on given network interface on customer edge site. It is template configuration that can be applied on a set of sites represented by list of sites...

## Example Usage

```hcl
data "f5_distributed_cloud_bgp" "example" {
  name      = "example-bgp"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
