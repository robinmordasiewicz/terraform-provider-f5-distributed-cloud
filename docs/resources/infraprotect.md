---
page_title: "f5xc_infraprotect Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  APIs to get monitoring data for infraprotect.
---

# f5xc_infraprotect (Resource)

APIs to get monitoring data for infraprotect.

## Example Usage

```hcl
resource "f5xc_infraprotect" "example" {
  name        = "example-infraprotect"
  namespace   = "system"
  description = "Example Infraprotect resource"
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

Infraprotect can be imported using the namespace and name:

```shell
terraform import f5xc_infraprotect.example namespace/name
```
