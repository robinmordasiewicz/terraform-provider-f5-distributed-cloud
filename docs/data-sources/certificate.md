---
page_title: "f5xc_certificate Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certificate represents a client or server certificate.
---

# f5xc_certificate (Data Source)

Certificate represents a client or server certificate.

## Example Usage

```hcl
data "f5xc_certificate" "example" {
  name      = "example-certificate"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
