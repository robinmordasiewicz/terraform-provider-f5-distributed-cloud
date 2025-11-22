---
page_title: "f5xc_rbac_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A rbac_policy object consists of list of rbac policy rules that when assigned to a user via Role object, it controls access of an user to list of APIs defined under the API Group name. Each rule un...
---

# f5xc_rbac_policy (Data Source)

A rbac_policy object consists of list of rbac policy rules that when assigned to a user via Role object, it controls access of an user to list of APIs defined under the API Group name. Each rule un...

## Example Usage

```hcl
data "f5xc_rbac_policy" "example" {
  name      = "example-rbac_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
