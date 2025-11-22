---
page_title: "f5xc_enhanced_firewall_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Enhanced Firewall Policy defined firewall rules applied in the site Enhanced Firewall Policy defines rules in the form or 5 tuple, <source, source-port, destination, destination-port, protocol>  Wh...
---

# f5xc_enhanced_firewall_policy (Data Source)

Enhanced Firewall Policy defined firewall rules applied in the site Enhanced Firewall Policy defines rules in the form or 5 tuple, <source, source-port, destination, destination-port, protocol>  Wh...

## Example Usage

```hcl
data "f5xc_enhanced_firewall_policy" "example" {
  name      = "example-enhanced_firewall_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
