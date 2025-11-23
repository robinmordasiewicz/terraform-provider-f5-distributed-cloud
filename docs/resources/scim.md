---
page_title: "f5_distributed_cloud_scim Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  This schema specification details Volterra's support for SCIM protocol. Admin can use SCIM feature on top of SSO to enable automated provisioning of user and user groups from external identity prov...
---

# f5_distributed_cloud_scim (Resource)

This schema specification details Volterra's support for SCIM protocol. Admin can use SCIM feature on top of SSO to enable automated provisioning of user and user groups from external identity prov...

## Example Usage

```hcl
resource "f5_distributed_cloud_scim" "example" {
  name        = "example-scim"
  namespace   = "system"
  description = "Example Scim resource"
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

Scim can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_scim.example namespace/name
```
