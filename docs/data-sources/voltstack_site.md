---
page_title: "f5xc_voltstack_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  App Stack site defines a required parameters that can be used in CRUD, to create and manage an App Stack site.
---

# f5xc_voltstack_site (Data Source)

App Stack site defines a required parameters that can be used in CRUD, to create and manage an App Stack site.

## Example Usage

```hcl
data "f5xc_voltstack_site" "example" {
  name      = "example-voltstack_site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
