---
page_title: "f5_distributed_cloud_network_policy_rule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Policy Rule is applied to given remote endpoints to and from a given local endpoint and is a terminal rule. Local Endpoint is picked from the network policy to which this rule has been atta...
---

# f5_distributed_cloud_network_policy_rule (Data Source)

Network Policy Rule is applied to given remote endpoints to and from a given local endpoint and is a terminal rule. Local Endpoint is picked from the network policy to which this rule has been atta...

## Example Usage

```hcl
data "f5_distributed_cloud_network_policy_rule" "example" {
  name      = "example-network_policy_rule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
