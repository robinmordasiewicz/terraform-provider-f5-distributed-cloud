---
page_title: "f5xc_infraprotect_deny_list_rule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Deny List Rule information
---

# f5xc_infraprotect_deny_list_rule (Data Source)

DDoS transit Deny List Rule information

## Example Usage

```hcl
data "f5xc_infraprotect_deny_list_rule" "example" {
  name      = "example-infraprotect_deny_list_rule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
