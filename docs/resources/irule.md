---
page_title: "f5xc_irule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  iRule object defines the iRule that can be used in CRUD, to create and manage iRule for manipulating the application traffic.
---

# f5xc_irule (Resource)

iRule object defines the iRule that can be used in CRUD, to create and manage iRule for manipulating the application traffic.

## Example Usage

```hcl
resource "f5xc_irule" "example" {
  name        = "example-irule"
  namespace   = "system"
  description = "Example Irule resource"
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

Irule can be imported using the namespace and name:

```shell
terraform import f5xc_irule.example namespace/name
```
