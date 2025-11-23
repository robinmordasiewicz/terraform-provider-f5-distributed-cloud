---
page_title: "f5_distributed_cloud_udp_loadbalancer Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  UDP load balancer view defines a required parameters that can be used in CRUD, to create and manage UDP load balancer. It can be used to create UDP load balancer.  View will create following child ...
---

# f5_distributed_cloud_udp_loadbalancer (Data Source)

UDP load balancer view defines a required parameters that can be used in CRUD, to create and manage UDP load balancer. It can be used to create UDP load balancer.  View will create following child ...

## Example Usage

```hcl
data "f5_distributed_cloud_udp_loadbalancer" "example" {
  name      = "example-udp_loadbalancer"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
