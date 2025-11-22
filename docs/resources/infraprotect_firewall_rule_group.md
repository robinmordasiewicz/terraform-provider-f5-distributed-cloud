---
page_title: "f5xc_infraprotect_firewall_rule_group Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Firewall Rule Group
---

# f5xc_infraprotect_firewall_rule_group (Resource)

DDoS transit Firewall Rule Group

## Example Usage

```hcl
resource "f5xc_infraprotect_firewall_rule_group" "example" {
  name        = "example-infraprotect_firewall_rule_group"
  namespace   = "system"
  description = "Example InfraprotectFirewallRuleGroup resource"
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

InfraprotectFirewallRuleGroup can be imported using the namespace and name:

```shell
terraform import f5xc_infraprotect_firewall_rule_group.example namespace/name
```
