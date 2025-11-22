---
page_title: "f5xc_secret_management_access Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  secret_management_access object is used to define configuration on how to connect to a secret management backend. If secret backend is not F5XC Secret Management System, this objects needs to be co...
---

# f5xc_secret_management_access (Resource)

secret_management_access object is used to define configuration on how to connect to a secret management backend. If secret backend is not F5XC Secret Management System, this objects needs to be co...

## Example Usage

```hcl
resource "f5xc_secret_management_access" "example" {
  name        = "example-secret_management_access"
  namespace   = "system"
  description = "Example SecretManagementAccess resource"
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

SecretManagementAccess can be imported using the namespace and name:

```shell
terraform import f5xc_secret_management_access.example namespace/name
```
