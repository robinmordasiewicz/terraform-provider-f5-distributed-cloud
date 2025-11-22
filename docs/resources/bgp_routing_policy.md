---
page_title: "f5xc_bgp_routing_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BGP Routing Policy is a list of rules, which contains match criteria and action to be taken upon match. This helps in filtering routes which are exported or imported by BGP peers
---

# f5xc_bgp_routing_policy (Resource)

BGP Routing Policy is a list of rules, which contains match criteria and action to be taken upon match. This helps in filtering routes which are exported or imported by BGP peers

## Example Usage

```hcl
resource "f5xc_bgp_routing_policy" "example" {
  name        = "example-bgp_routing_policy"
  namespace   = "system"
  description = "Example BGPRoutingPolicy resource"
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

BGPRoutingPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_bgp_routing_policy.example namespace/name
```
