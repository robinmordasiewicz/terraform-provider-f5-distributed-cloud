---
page_title: "f5xc_enhanced_firewall_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Enhanced Firewall Policy defined firewall rules applied in the site Enhanced Firewall Policy defines rules in the form or 5 tuple, <source, source-port, destination, destination-port, protocol>  Wh...
---

# f5xc_enhanced_firewall_policy (Resource)

Enhanced Firewall Policy defined firewall rules applied in the site Enhanced Firewall Policy defines rules in the form or 5 tuple, <source, source-port, destination, destination-port, protocol>  Wh...

## Example Usage

```hcl
resource "f5xc_enhanced_firewall_policy" "example" {
  name        = "example-enhanced_firewall_policy"
  namespace   = "system"
  description = "Example EnhancedFirewallPolicy resource"
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

EnhancedFirewallPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_enhanced_firewall_policy.example namespace/name
```
