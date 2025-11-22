---
page_title: "f5xc_infraprotect_internet_prefix_advertisement Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Internet Prefix information
---

# f5xc_infraprotect_internet_prefix_advertisement (Resource)

DDoS transit Internet Prefix information

## Example Usage

```hcl
resource "f5xc_infraprotect_internet_prefix_advertisement" "example" {
  name        = "example-infraprotect_internet_prefix_advertisement"
  namespace   = "system"
  description = "Example InfraprotectInternetPrefixAdvertisement resource"
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

InfraprotectInternetPrefixAdvertisement can be imported using the namespace and name:

```shell
terraform import f5xc_infraprotect_internet_prefix_advertisement.example namespace/name
```
