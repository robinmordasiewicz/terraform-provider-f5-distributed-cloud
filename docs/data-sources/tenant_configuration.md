---
page_title: "f5xc_tenant_configuration Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Tenant configuration consists of three main parts: - Basic Configuration - Brute Force Detection Settings - Password Policy Basic configuration contains general parameters which can be adjusted wit...
---

# f5xc_tenant_configuration (Data Source)

Tenant configuration consists of three main parts: - Basic Configuration - Brute Force Detection Settings - Password Policy Basic configuration contains general parameters which can be adjusted wit...

## Example Usage

```hcl
data "f5xc_tenant_configuration" "example" {
  name      = "example-tenant_configuration"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
