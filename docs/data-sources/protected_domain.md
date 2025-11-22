---
page_title: "f5xc_protected_domain Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Domain to Protect Object defines which domains will be protected by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...
---

# f5xc_protected_domain (Data Source)

Domain to Protect Object defines which domains will be protected by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...

## Example Usage

```hcl
data "f5xc_protected_domain" "example" {
  name      = "example-protected_domain"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
