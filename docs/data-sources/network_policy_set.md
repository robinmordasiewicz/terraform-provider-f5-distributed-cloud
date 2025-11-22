---
page_title: "f5xc_network_policy_set Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network policy set implements L3/L4 stateful firewall. It is a list of one or more Network policy references and are applied sequentially in order specified in the list.  Network policy set can be ...
---

# f5xc_network_policy_set (Data Source)

Network policy set implements L3/L4 stateful firewall. It is a list of one or more Network policy references and are applied sequentially in order specified in the list.  Network policy set can be ...

## Example Usage

```hcl
data "f5xc_network_policy_set" "example" {
  name      = "example-network_policy_set"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
