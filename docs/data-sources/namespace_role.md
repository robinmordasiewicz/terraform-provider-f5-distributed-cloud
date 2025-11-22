---
page_title: "f5xc_namespace_role Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Namespace role defines a user's role in a namespace. Namespace role links a user with a role namespace. Using this object one can assign/remove a role to a user in a namespace. Namespace roles are ...
---

# f5xc_namespace_role (Data Source)

Namespace role defines a user's role in a namespace. Namespace role links a user with a role namespace. Using this object one can assign/remove a role to a user in a namespace. Namespace roles are ...

## Example Usage

```hcl
data "f5xc_namespace_role" "example" {
  name      = "example-namespace_role"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
