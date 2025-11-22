---
page_title: "f5xc_network_firewall Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Firewall is applicable when referred to by a Fleet. The Network Firewall will be applied to all sites that are member of the referring Fleet.  It is applied on all virtual networks that are...
---

# f5xc_network_firewall (Resource)

Network Firewall is applicable when referred to by a Fleet. The Network Firewall will be applied to all sites that are member of the referring Fleet.  It is applied on all virtual networks that are...

## Example Usage

```hcl
resource "f5xc_network_firewall" "example" {
  name        = "example-network_firewall"
  namespace   = "system"
  description = "Example NetworkFirewall resource"
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

NetworkFirewall can be imported using the namespace and name:

```shell
terraform import f5xc_network_firewall.example namespace/name
```
