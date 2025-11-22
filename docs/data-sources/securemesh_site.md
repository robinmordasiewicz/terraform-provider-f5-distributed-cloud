---
page_title: "f5xc_securemesh_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Secure Mesh site defines a required parameters that can be used in CRUD, to create and manage an Secure Mesh site.
---

# f5xc_securemesh_site (Data Source)

Secure Mesh site defines a required parameters that can be used in CRUD, to create and manage an Secure Mesh site.

## Example Usage

```hcl
data "f5xc_securemesh_site" "example" {
  name      = "example-securemesh_site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
