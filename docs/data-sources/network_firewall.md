---
page_title: "f5_distributed_cloud_network_firewall Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Firewall is applicable when referred to by a Fleet. The Network Firewall will be applied to all sites that are member of the referring Fleet.  It is applied on all virtual networks that are...
---

# f5_distributed_cloud_network_firewall (Data Source)

Network Firewall is applicable when referred to by a Fleet. The Network Firewall will be applied to all sites that are member of the referring Fleet.  It is applied on all virtual networks that are...

## Example Usage

```hcl
data "f5_distributed_cloud_network_firewall" "example" {
  name      = "example-network_firewall"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
