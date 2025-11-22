---
page_title: "f5xc_tcp_loadbalancer Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  TCP load balancer view defines a required parameters that can be used in CRUD, to create and manage TCP load balancer. It can be used to create TCP load balancer and TCP load balancer with SNI.  Vi...
---

# f5xc_tcp_loadbalancer (Resource)

TCP load balancer view defines a required parameters that can be used in CRUD, to create and manage TCP load balancer. It can be used to create TCP load balancer and TCP load balancer with SNI.  Vi...

## Example Usage

```hcl
resource "f5xc_tcp_loadbalancer" "example" {
  name        = "example-tcp_loadbalancer"
  namespace   = "system"
  description = "Example TCPLoadbalancer resource"
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

TCPLoadbalancer can be imported using the namespace and name:

```shell
terraform import f5xc_tcp_loadbalancer.example namespace/name
```
