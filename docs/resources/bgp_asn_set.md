---
page_title: "f5xc_bgp_asn_set Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or blacklists for use in network policy or service policy. The source or destination public IP address f...
---

# f5xc_bgp_asn_set (Resource)

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or blacklists for use in network policy or service policy. The source or destination public IP address f...

## Example Usage

```hcl
resource "f5xc_bgp_asn_set" "example" {
  name        = "example-bgp_asn_set"
  namespace   = "system"
  description = "Example BGPAsnSet resource"
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

BGPAsnSet can be imported using the namespace and name:

```shell
terraform import f5xc_bgp_asn_set.example namespace/name
```
