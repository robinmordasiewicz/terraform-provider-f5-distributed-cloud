---
page_title: "f5xc_network_interface Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Interface object represents the configuration of a network device in a fleet of F5XC Customer Edge sites. The following properties can be configured in this object:     IP address allocatio...
---

# f5xc_network_interface (Resource)

Network Interface object represents the configuration of a network device in a fleet of F5XC Customer Edge sites. The following properties can be configured in this object:     IP address allocatio...

## Example Usage

```hcl
resource "f5xc_network_interface" "example" {
  name        = "example-network_interface"
  namespace   = "system"
  description = "Example NetworkInterface resource"
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

NetworkInterface can be imported using the namespace and name:

```shell
terraform import f5xc_network_interface.example namespace/name
```
