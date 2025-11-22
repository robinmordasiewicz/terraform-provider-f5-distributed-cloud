---
page_title: "f5xc_network_interface Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Interface object represents the configuration of a network device in a fleet of F5XC Customer Edge sites. The following properties can be configured in this object:     IP address allocatio...
---

# f5xc_network_interface (Data Source)

Network Interface object represents the configuration of a network device in a fleet of F5XC Customer Edge sites. The following properties can be configured in this object:     IP address allocatio...

## Example Usage

```hcl
data "f5xc_network_interface" "example" {
  name      = "example-network_interface"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
