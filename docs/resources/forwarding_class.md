---
page_title: "f5_distributed_cloud_forwarding_class Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  In Policy Based Routing(forwarding) (PBR) PBR policy can select Forwarding Class object as action When match condition is satisfied(true). Forwarding class determines three things:    Ordered list ...
---

# f5_distributed_cloud_forwarding_class (Resource)

In Policy Based Routing(forwarding) (PBR) PBR policy can select Forwarding Class object as action When match condition is satisfied(true). Forwarding class determines three things:    Ordered list ...

## Example Usage

```hcl
resource "f5_distributed_cloud_forwarding_class" "example" {
  name        = "example-forwarding_class"
  namespace   = "system"
  description = "Example ForwardingClass resource"
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

ForwardingClass can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_forwarding_class.example namespace/name
```
