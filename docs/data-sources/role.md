---
page_title: "f5xc_role Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Defines the role the user has in a namespace. There are two kinds of roles: * user defines roles - each tenant can define a set of their roles suiting their needs. * built-in roles - volterra defin...
---

# f5xc_role (Data Source)

Defines the role the user has in a namespace. There are two kinds of roles: * user defines roles - each tenant can define a set of their roles suiting their needs. * built-in roles - volterra defin...

## Example Usage

```hcl
data "f5xc_role" "example" {
  name      = "example-role"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
