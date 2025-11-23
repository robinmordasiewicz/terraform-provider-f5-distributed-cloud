---
page_title: "f5_distributed_cloud_infraprotect_asn Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit ASN information
---

# f5_distributed_cloud_infraprotect_asn (Resource)

DDoS transit ASN information

## Example Usage

```hcl
resource "f5_distributed_cloud_infraprotect_asn" "example" {
  name        = "example-infraprotect_asn"
  namespace   = "system"
  description = "Example InfraprotectAsn resource"
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

InfraprotectAsn can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_infraprotect_asn.example namespace/name
```
