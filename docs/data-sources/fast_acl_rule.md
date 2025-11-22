---
page_title: "f5xc_fast_acl_rule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Fast ACL rule Rule consists of following:   1. Protocol to match from packet   2. List of source IP prefixes or a reference to IP prefix set, to which      source IP in packet should belong to.   3...
---

# f5xc_fast_acl_rule (Data Source)

Fast ACL rule Rule consists of following:   1. Protocol to match from packet   2. List of source IP prefixes or a reference to IP prefix set, to which      source IP in packet should belong to.   3...

## Example Usage

```hcl
data "f5xc_fast_acl_rule" "example" {
  name      = "example-fast_acl_rule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
