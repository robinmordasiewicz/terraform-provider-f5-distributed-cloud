---
page_title: "f5_distributed_cloud_infraprotect_asn_prefix Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Prefix Information
---

# f5_distributed_cloud_infraprotect_asn_prefix (Resource)

DDoS transit Prefix Information

## Example Usage

```hcl
resource "f5_distributed_cloud_infraprotect_asn_prefix" "example" {
  name        = "example-infraprotect_asn_prefix"
  namespace   = "system"
  description = "Example InfraprotectAsnPrefix resource"
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

InfraprotectAsnPrefix can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_infraprotect_asn_prefix.example namespace/name
```
