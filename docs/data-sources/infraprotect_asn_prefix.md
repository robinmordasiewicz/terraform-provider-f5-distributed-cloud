---
page_title: "f5xc_infraprotect_asn_prefix Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Prefix Information
---

# f5xc_infraprotect_asn_prefix (Data Source)

DDoS transit Prefix Information

## Example Usage

```hcl
data "f5xc_infraprotect_asn_prefix" "example" {
  name      = "example-infraprotect_asn_prefix"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
