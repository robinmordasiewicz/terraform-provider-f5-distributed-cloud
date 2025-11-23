---
page_title: "f5_distributed_cloud_tenant_configuration Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Tenant configuration consists of three main parts: - Basic Configuration - Brute Force Detection Settings - Password Policy Basic configuration contains general parameters which can be adjusted wit...
---

# f5_distributed_cloud_tenant_configuration (Resource)

Tenant configuration consists of three main parts: - Basic Configuration - Brute Force Detection Settings - Password Policy Basic configuration contains general parameters which can be adjusted wit...

## Example Usage

```hcl
resource "f5_distributed_cloud_tenant_configuration" "example" {
  name        = "example-tenant_configuration"
  namespace   = "system"
  description = "Example TenantConfiguration resource"
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

TenantConfiguration can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_tenant_configuration.example namespace/name
```
