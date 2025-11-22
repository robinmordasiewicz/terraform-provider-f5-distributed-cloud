---
page_title: "f5xc_app_security Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  API to create API endpoint protection rule suggestion from App Security Monitoring pages
---

# f5xc_app_security (Data Source)

API to create API endpoint protection rule suggestion from App Security Monitoring pages

## Example Usage

```hcl
data "f5xc_app_security" "example" {
  name      = "example-app_security"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
