---
page_title: "f5_distributed_cloud_secret_management_access Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  secret_management_access object is used to define configuration on how to connect to a secret management backend. If secret backend is not F5XC Secret Management System, this objects needs to be co...
---

# f5_distributed_cloud_secret_management_access (Data Source)

secret_management_access object is used to define configuration on how to connect to a secret management backend. If secret backend is not F5XC Secret Management System, this objects needs to be co...

## Example Usage

```hcl
data "f5_distributed_cloud_secret_management_access" "example" {
  name      = "example-secret_management_access"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
