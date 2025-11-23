---
page_title: "f5_distributed_cloud_bigip_virtual_server Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BIG-IP virtual server view repesents the internal virtual host corresponding to the virtual-servers discovered from BIG-IPs It exposes parameters to enable API discovery and other WAAP security fea...
---

# f5_distributed_cloud_bigip_virtual_server (Data Source)

BIG-IP virtual server view repesents the internal virtual host corresponding to the virtual-servers discovered from BIG-IPs It exposes parameters to enable API discovery and other WAAP security fea...

## Example Usage

```hcl
data "f5_distributed_cloud_bigip_virtual_server" "example" {
  name      = "example-bigip_virtual_server"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
