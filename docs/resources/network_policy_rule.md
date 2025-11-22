---
page_title: "f5xc_network_policy_rule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Policy Rule is applied to given remote endpoints to and from a given local endpoint and is a terminal rule. Local Endpoint is picked from the network policy to which this rule has been atta...
---

# f5xc_network_policy_rule (Resource)

Network Policy Rule is applied to given remote endpoints to and from a given local endpoint and is a terminal rule. Local Endpoint is picked from the network policy to which this rule has been atta...

## Example Usage

```hcl
resource "f5xc_network_policy_rule" "example" {
  name        = "example-network_policy_rule"
  namespace   = "system"
  description = "Example NetworkPolicyRule resource"
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

NetworkPolicyRule can be imported using the namespace and name:

```shell
terraform import f5xc_network_policy_rule.example namespace/name
```
