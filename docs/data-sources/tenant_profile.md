---
page_title: "f5xc_tenant_profile Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Tenant profile objects are required for creating child tenant using Child Tenant Manager as part of MSP. Tenant Profile is the template which defines the child tenant configuration properties e.g.,...
---

# f5xc_tenant_profile (Data Source)

Tenant profile objects are required for creating child tenant using Child Tenant Manager as part of MSP. Tenant Profile is the template which defines the child tenant configuration properties e.g.,...

## Example Usage

```hcl
data "f5xc_tenant_profile" "example" {
  name      = "example-tenant_profile"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
