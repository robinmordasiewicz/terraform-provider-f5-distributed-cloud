---
page_title: "f5xc_protected_application Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configures application protected by Bot Defense
---

# f5xc_protected_application (Resource)

Configures application protected by Bot Defense

## Example Usage

```hcl
resource "f5xc_protected_application" "example" {
  name        = "example-protected_application"
  namespace   = "system"
  description = "Example ProtectedApplication resource"
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

ProtectedApplication can be imported using the namespace and name:

```shell
terraform import f5xc_protected_application.example namespace/name
```
