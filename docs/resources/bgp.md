---
page_title: "f5_distributed_cloud_bgp Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BGP object represents configuration of bgp protocol on given network interface on customer edge site. It is template configuration that can be applied on a set of sites represented by list of sites...
---

# f5_distributed_cloud_bgp (Resource)

BGP object represents configuration of bgp protocol on given network interface on customer edge site. It is template configuration that can be applied on a set of sites represented by list of sites...

## Example Usage

```hcl
resource "f5_distributed_cloud_bgp" "example" {
  name        = "example-bgp"
  namespace   = "system"
  description = "Example BGP resource"
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

BGP can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_bgp.example namespace/name
```
