---
page_title: "f5_distributed_cloud_ip_prefix_set Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  An ip prefix set contains an unordered list of IP prefixes. It can can be used to whitelist or blacklist specific IP prefixes via network policy or service policy.
---

# f5_distributed_cloud_ip_prefix_set (Data Source)

An ip prefix set contains an unordered list of IP prefixes. It can can be used to whitelist or blacklist specific IP prefixes via network policy or service policy.

## Example Usage

```hcl
data "f5_distributed_cloud_ip_prefix_set" "example" {
  name      = "example-ip_prefix_set"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
