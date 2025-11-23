---
page_title: "f5_distributed_cloud_infraprotect_firewall_ruleset Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Firewall Ruleset
---

# f5_distributed_cloud_infraprotect_firewall_ruleset (Data Source)

DDoS transit Firewall Ruleset

## Example Usage

```hcl
data "f5_distributed_cloud_infraprotect_firewall_ruleset" "example" {
  name      = "example-infraprotect_firewall_ruleset"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
