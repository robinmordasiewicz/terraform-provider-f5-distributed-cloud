---
page_title: "f5xc_udp_loadbalancer Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  UDP load balancer view defines a required parameters that can be used in CRUD, to create and manage UDP load balancer. It can be used to create UDP load balancer.  View will create following child ...
---

# f5xc_udp_loadbalancer (Resource)

UDP load balancer view defines a required parameters that can be used in CRUD, to create and manage UDP load balancer. It can be used to create UDP load balancer.  View will create following child ...

## Example Usage

```hcl
resource "f5xc_udp_loadbalancer" "example" {
  name        = "example-udp_loadbalancer"
  namespace   = "system"
  description = "Example UDPLoadbalancer resource"
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

UDPLoadbalancer can be imported using the namespace and name:

```shell
terraform import f5xc_udp_loadbalancer.example namespace/name
```
