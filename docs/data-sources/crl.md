---
page_title: "f5_distributed_cloud_crl Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certificate Revocation List(CRL)  It contains information about CRL server and how to download the CRL file CRL file is used to validate the certificate presented to check whether it is revoked or ...
---

# f5_distributed_cloud_crl (Data Source)

Certificate Revocation List(CRL)  It contains information about CRL server and how to download the CRL file CRL file is used to validate the certificate presented to check whether it is revoked or ...

## Example Usage

```hcl
data "f5_distributed_cloud_crl" "example" {
  name      = "example-crl"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
