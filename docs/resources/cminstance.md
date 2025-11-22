---
page_title: "f5xc_cminstance Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  cminsatnce object can be used to enable connectivity between ce site and bigip central manager.
---

# f5xc_cminstance (Resource)

cminsatnce object can be used to enable connectivity between ce site and bigip central manager.

## Example Usage

```hcl
resource "f5xc_cminstance" "example" {
  name        = "example-cminstance"
  namespace   = "system"
  description = "Example Cminstance resource"
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

Cminstance can be imported using the namespace and name:

```shell
terraform import f5xc_cminstance.example namespace/name
```
