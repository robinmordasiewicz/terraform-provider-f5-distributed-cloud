---
page_title: "f5xc_infraprotect_internet_prefix_advertisement Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Internet Prefix information
---

# f5xc_infraprotect_internet_prefix_advertisement (Data Source)

DDoS transit Internet Prefix information

## Example Usage

```hcl
data "f5xc_infraprotect_internet_prefix_advertisement" "example" {
  name      = "example-infraprotect_internet_prefix_advertisement"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
