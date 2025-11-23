---
page_title: "f5_distributed_cloud_certificate_chain Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certificate chain is list of certificates used to establish chain of trust from server or client certificate to trusted CA root certificates.
---

# f5_distributed_cloud_certificate_chain (Resource)

Certificate chain is list of certificates used to establish chain of trust from server or client certificate to trusted CA root certificates.

## Example Usage

```hcl
resource "f5_distributed_cloud_certificate_chain" "example" {
  name        = "example-certificate_chain"
  namespace   = "system"
  description = "Example CertificateChain resource"
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

CertificateChain can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_certificate_chain.example namespace/name
```
