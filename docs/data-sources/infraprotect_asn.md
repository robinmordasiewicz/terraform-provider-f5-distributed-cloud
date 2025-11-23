---
page_title: "f5_distributed_cloud_infraprotect_asn Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit ASN information
---

# f5_distributed_cloud_infraprotect_asn (Data Source)

DDoS transit ASN information

## Example Usage

```hcl
data "f5_distributed_cloud_infraprotect_asn" "example" {
  name      = "example-infraprotect_asn"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
