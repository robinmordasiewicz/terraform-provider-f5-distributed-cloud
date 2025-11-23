---
page_title: "f5_distributed_cloud_protocol_policer Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Protocol policer has set or network protocol fields and flags to be match on a layer 4 packet and corresponding rate limit to be applied, this would be useful in specifying * Ratelimiting TCP sessi...
---

# f5_distributed_cloud_protocol_policer (Resource)

Protocol policer has set or network protocol fields and flags to be match on a layer 4 packet and corresponding rate limit to be applied, this would be useful in specifying * Ratelimiting TCP sessi...

## Example Usage

```hcl
resource "f5_distributed_cloud_protocol_policer" "example" {
  name        = "example-protocol_policer"
  namespace   = "system"
  description = "Example ProtocolPolicer resource"
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

ProtocolPolicer can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_protocol_policer.example namespace/name
```
