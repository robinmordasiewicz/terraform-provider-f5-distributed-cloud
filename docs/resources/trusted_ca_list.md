---
page_title: "f5xc_trusted_ca_list Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A Root CA Certificate represents list of trusted root CAs
---

# f5xc_trusted_ca_list (Resource)

A Root CA Certificate represents list of trusted root CAs

## Example Usage

```hcl
resource "f5xc_trusted_ca_list" "example" {
  name        = "example-trusted_ca_list"
  namespace   = "system"
  description = "Example TrustedCAList resource"
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

TrustedCAList can be imported using the namespace and name:

```shell
terraform import f5xc_trusted_ca_list.example namespace/name
```
