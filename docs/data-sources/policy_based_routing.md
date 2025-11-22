---
page_title: "f5xc_policy_based_routing Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Policy based routing is used to control how different classes of traffic is forwarded and QOS is applied over WAN interfaces in SDWAN scenarios.
---

# f5xc_policy_based_routing (Data Source)

Policy based routing is used to control how different classes of traffic is forwarded and QOS is applied over WAN interfaces in SDWAN scenarios.

## Example Usage

```hcl
data "f5xc_policy_based_routing" "example" {
  name      = "example-policy_based_routing"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
