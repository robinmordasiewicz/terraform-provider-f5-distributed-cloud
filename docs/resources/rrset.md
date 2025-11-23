---
page_title: "f5_distributed_cloud_rrset Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  x-required APIs to create, update or delete individual records of a DNS zone without having to send the whole DNS zone information.
---

# f5_distributed_cloud_rrset (Resource)

x-required APIs to create, update or delete individual records of a DNS zone without having to send the whole DNS zone information.

## Example Usage

```hcl
resource "f5_distributed_cloud_rrset" "example" {
  name        = "example-rrset"
  namespace   = "system"
  description = "Example Rrset resource"
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

Rrset can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_rrset.example namespace/name
```
