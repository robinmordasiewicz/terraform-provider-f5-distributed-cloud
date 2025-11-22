---
page_title: "f5xc_bgp_asn_set Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or blacklists for use in network policy or service policy. The source or destination public IP address f...
---

# f5xc_bgp_asn_set (Data Source)

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or blacklists for use in network policy or service policy. The source or destination public IP address f...

## Example Usage

```hcl
data "f5xc_bgp_asn_set" "example" {
  name      = "example-bgp_asn_set"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
