---
page_title: "f5xc_tpm_api_key Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  TPM API Keys are used by TPM provisioning tool during manufacturing to call in to F5XC TPM provisioning service to generate EK certificates and provision them in the Volterra's Customer Edge(CE) de...
---

# f5xc_tpm_api_key (Data Source)

TPM API Keys are used by TPM provisioning tool during manufacturing to call in to F5XC TPM provisioning service to generate EK certificates and provision them in the Volterra's Customer Edge(CE) de...

## Example Usage

```hcl
data "f5xc_tpm_api_key" "example" {
  name      = "example-tpm_api_key"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
