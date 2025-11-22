---
page_title: "f5xc_policer Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  * Policer objects enforces traffic rate limits * network_policy_rule and fast_acl_rule can refer to a policer. Packets   matching those rules will be subjected to the traffic contract specified in ...
---

# f5xc_policer (Data Source)

* Policer objects enforces traffic rate limits * network_policy_rule and fast_acl_rule can refer to a policer. Packets   matching those rules will be subjected to the traffic contract specified in ...

## Example Usage

```hcl
data "f5xc_policer" "example" {
  name      = "example-policer"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
