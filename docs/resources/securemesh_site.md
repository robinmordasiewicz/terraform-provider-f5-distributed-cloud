---
page_title: "f5xc_securemesh_site Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Secure Mesh site defines a required parameters that can be used in CRUD, to create and manage an Secure Mesh site.
---

# f5xc_securemesh_site (Resource)

Secure Mesh site defines a required parameters that can be used in CRUD, to create and manage an Secure Mesh site.

## Example Usage

```hcl
resource "f5xc_securemesh_site" "example" {
  name        = "example-securemesh_site"
  namespace   = "system"
  description = "Example SecuremeshSite resource"
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

SecuremeshSite can be imported using the namespace and name:

```shell
terraform import f5xc_securemesh_site.example namespace/name
```
