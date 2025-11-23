---
page_title: "f5_distributed_cloud_fast_acl_rule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Fast ACL rule Rule consists of following:   1. Protocol to match from packet   2. List of source IP prefixes or a reference to IP prefix set, to which      source IP in packet should belong to.   3...
---

# f5_distributed_cloud_fast_acl_rule (Resource)

Fast ACL rule Rule consists of following:   1. Protocol to match from packet   2. List of source IP prefixes or a reference to IP prefix set, to which      source IP in packet should belong to.   3...

## Example Usage

```hcl
resource "f5_distributed_cloud_fast_acl_rule" "example" {
  name        = "example-fast_acl_rule"
  namespace   = "system"
  description = "Example FastACLRule resource"
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

FastACLRule can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_fast_acl_rule.example namespace/name
```
