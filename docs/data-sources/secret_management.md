---
page_title: "f5_distributed_cloud_secret_management Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  F5XC Secret Management service serves APIs for information required for offline secret encryption such as getting the public key and getting the secret policy document.
---

# f5_distributed_cloud_secret_management (Data Source)

F5XC Secret Management service serves APIs for information required for offline secret encryption such as getting the public key and getting the secret policy document.

## Example Usage

```hcl
data "f5_distributed_cloud_secret_management" "example" {
  name      = "example-secret_management"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
