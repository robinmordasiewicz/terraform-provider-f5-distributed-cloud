---
page_title: "f5xc_child_tenant_manager Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Child Tenant Manager uses Tenant Profile template to create child tenant and store its object reference.  Any look up for child tenant should be done from child tenant manager since it stores objec...
---

# f5xc_child_tenant_manager (Resource)

Child Tenant Manager uses Tenant Profile template to create child tenant and store its object reference.  Any look up for child tenant should be done from child tenant manager since it stores objec...

## Example Usage

```hcl
resource "f5xc_child_tenant_manager" "example" {
  name        = "example-child_tenant_manager"
  namespace   = "system"
  description = "Example ChildTenantManager resource"
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

ChildTenantManager can be imported using the namespace and name:

```shell
terraform import f5xc_child_tenant_manager.example namespace/name
```
