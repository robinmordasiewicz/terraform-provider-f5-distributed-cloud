---
page_title: "f5xc_protected_domain Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Domain to Protect Object defines which domains will be protected by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...
---

# f5xc_protected_domain (Resource)

Domain to Protect Object defines which domains will be protected by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...

## Example Usage

```hcl
resource "f5xc_protected_domain" "example" {
  name        = "example-protected_domain"
  namespace   = "system"
  description = "Example ProtectedDomain resource"
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

ProtectedDomain can be imported using the namespace and name:

```shell
terraform import f5xc_protected_domain.example namespace/name
```
