---
page_title: "f5xc_crl Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certificate Revocation List(CRL)  It contains information about CRL server and how to download the CRL file CRL file is used to validate the certificate presented to check whether it is revoked or ...
---

# f5xc_crl (Resource)

Certificate Revocation List(CRL)  It contains information about CRL server and how to download the CRL file CRL file is used to validate the certificate presented to check whether it is revoked or ...

## Example Usage

```hcl
resource "f5xc_crl" "example" {
  name        = "example-crl"
  namespace   = "system"
  description = "Example CRL resource"
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

CRL can be imported using the namespace and name:

```shell
terraform import f5xc_crl.example namespace/name
```
