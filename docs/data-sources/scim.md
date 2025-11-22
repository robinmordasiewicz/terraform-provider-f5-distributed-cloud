---
page_title: "f5xc_scim Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  This schema specification details Volterra's support for SCIM protocol. Admin can use SCIM feature on top of SSO to enable automated provisioning of user and user groups from external identity prov...
---

# f5xc_scim (Data Source)

This schema specification details Volterra's support for SCIM protocol. Admin can use SCIM feature on top of SSO to enable automated provisioning of user and user groups from external identity prov...

## Example Usage

```hcl
data "f5xc_scim" "example" {
  name      = "example-scim"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
