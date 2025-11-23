---
page_title: "f5_distributed_cloud_bigip_irule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BIG-IP iRule Service manages iRule Library for customers
---

# f5_distributed_cloud_bigip_irule (Resource)

BIG-IP iRule Service manages iRule Library for customers

## Example Usage

```hcl
resource "f5_distributed_cloud_bigip_irule" "example" {
  name        = "example-bigip_irule"
  namespace   = "system"
  description = "Example BigipIrule resource"
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

BigipIrule can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_bigip_irule.example namespace/name
```
