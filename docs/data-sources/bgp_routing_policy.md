---
page_title: "f5_distributed_cloud_bgp_routing_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BGP Routing Policy is a list of rules, which contains match criteria and action to be taken upon match. This helps in filtering routes which are exported or imported by BGP peers
---

# f5_distributed_cloud_bgp_routing_policy (Data Source)

BGP Routing Policy is a list of rules, which contains match criteria and action to be taken upon match. This helps in filtering routes which are exported or imported by BGP peers

## Example Usage

```hcl
data "f5_distributed_cloud_bgp_routing_policy" "example" {
  name      = "example-bgp_routing_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
