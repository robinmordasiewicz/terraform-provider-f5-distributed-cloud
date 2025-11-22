---
page_title: "f5xc_certificate_chain Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certificate chain is list of certificates used to establish chain of trust from server or client certificate to trusted CA root certificates.
---

# f5xc_certificate_chain (Data Source)

Certificate chain is list of certificates used to establish chain of trust from server or client certificate to trusted CA root certificates.

## Example Usage

```hcl
data "f5xc_certificate_chain" "example" {
  name      = "example-certificate_chain"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
