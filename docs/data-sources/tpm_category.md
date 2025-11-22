---
page_title: "f5xc_tpm_category Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Category is a grouping of APIKeys, each category comes with its own SubCA for signing TPM EK, SRK and AK certificates
---

# f5xc_tpm_category (Data Source)

Category is a grouping of APIKeys, each category comes with its own SubCA for signing TPM EK, SRK and AK certificates

## Example Usage

```hcl
data "f5xc_tpm_category" "example" {
  name      = "example-tpm_category"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
