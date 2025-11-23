---
page_title: "f5_distributed_cloud_protocol_policer Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Protocol policer has set or network protocol fields and flags to be match on a layer 4 packet and corresponding rate limit to be applied, this would be useful in specifying * Ratelimiting TCP sessi...
---

# f5_distributed_cloud_protocol_policer (Data Source)

Protocol policer has set or network protocol fields and flags to be match on a layer 4 packet and corresponding rate limit to be applied, this would be useful in specifying * Ratelimiting TCP sessi...

## Example Usage

```hcl
data "f5_distributed_cloud_protocol_policer" "example" {
  name      = "example-protocol_policer"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
