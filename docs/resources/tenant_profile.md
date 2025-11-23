---
page_title: "f5_distributed_cloud_tenant_profile Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Tenant profile objects are required for creating child tenant using Child Tenant Manager as part of MSP. Tenant Profile is the template which defines the child tenant configuration properties e.g.,...
---

# f5_distributed_cloud_tenant_profile (Resource)

Tenant profile objects are required for creating child tenant using Child Tenant Manager as part of MSP. Tenant Profile is the template which defines the child tenant configuration properties e.g.,...

## Example Usage

```hcl
resource "f5_distributed_cloud_tenant_profile" "example" {
  name        = "example-tenant_profile"
  namespace   = "system"
  description = "Example TenantProfile resource"
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

TenantProfile can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_tenant_profile.example namespace/name
```
