---
page_title: "f5_distributed_cloud_tcp_loadbalancer Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  TCP load balancer view defines a required parameters that can be used in CRUD, to create and manage TCP load balancer. It can be used to create TCP load balancer and TCP load balancer with SNI.  Vi...
---

# f5_distributed_cloud_tcp_loadbalancer (Data Source)

TCP load balancer view defines a required parameters that can be used in CRUD, to create and manage TCP load balancer. It can be used to create TCP load balancer and TCP load balancer with SNI.  Vi...

## Example Usage

```hcl
data "f5_distributed_cloud_tcp_loadbalancer" "example" {
  name      = "example-tcp_loadbalancer"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
