---
page_title: "f5_distributed_cloud_infraprotect_firewall_rule_group Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Firewall Rule Group
---

# f5_distributed_cloud_infraprotect_firewall_rule_group (Data Source)

DDoS transit Firewall Rule Group

## Example Usage

```hcl
data "f5_distributed_cloud_infraprotect_firewall_rule_group" "example" {
  name      = "example-infraprotect_firewall_rule_group"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
