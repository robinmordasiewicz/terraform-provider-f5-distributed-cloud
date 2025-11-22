---
page_title: "f5xc_role Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Defines the role the user has in a namespace. There are two kinds of roles: * user defines roles - each tenant can define a set of their roles suiting their needs. * built-in roles - volterra defin...
---

# f5xc_role (Resource)

Defines the role the user has in a namespace. There are two kinds of roles: * user defines roles - each tenant can define a set of their roles suiting their needs. * built-in roles - volterra defin...

## Example Usage

```hcl
resource "f5xc_role" "example" {
  name        = "example-role"
  namespace   = "system"
  description = "Example Role resource"
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

Role can be imported using the namespace and name:

```shell
terraform import f5xc_role.example namespace/name
```
