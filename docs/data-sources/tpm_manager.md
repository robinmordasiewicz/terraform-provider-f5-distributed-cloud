---
page_title: "f5_distributed_cloud_tpm_manager Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  TPM Manager stores the required TPM management related data for the customer. The data includes APIKeys, Category, Root and SubCA's, and Client Certificates There are three types of Root CAs that a...
---

# f5_distributed_cloud_tpm_manager (Data Source)

TPM Manager stores the required TPM management related data for the customer. The data includes APIKeys, Category, Root and SubCA's, and Client Certificates There are three types of Root CAs that a...

## Example Usage

```hcl
data "f5_distributed_cloud_tpm_manager" "example" {
  name      = "example-tpm_manager"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
