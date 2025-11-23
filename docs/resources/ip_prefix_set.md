---
page_title: "f5_distributed_cloud_ip_prefix_set Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  An ip prefix set contains an unordered list of IP prefixes. It can can be used to whitelist or blacklist specific IP prefixes via network policy or service policy.
---

# f5_distributed_cloud_ip_prefix_set (Resource)

An ip prefix set contains an unordered list of IP prefixes. It can can be used to whitelist or blacklist specific IP prefixes via network policy or service policy.

## Example Usage

```hcl
resource "f5_distributed_cloud_ip_prefix_set" "example" {
  name        = "example-ip_prefix_set"
  namespace   = "system"
  description = "Example IPPrefixSet resource"
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

IPPrefixSet can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_ip_prefix_set.example namespace/name
```
