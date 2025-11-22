---
page_title: "f5xc_user_group Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Represents group for a given tenant
---

# f5xc_user_group (Resource)

Represents group for a given tenant

## Example Usage

```hcl
resource "f5xc_user_group" "example" {
  name        = "example-user_group"
  namespace   = "system"
  description = "Example UserGroup resource"
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

UserGroup can be imported using the namespace and name:

```shell
terraform import f5xc_user_group.example namespace/name
```
