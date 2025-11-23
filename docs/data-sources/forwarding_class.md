---
page_title: "f5_distributed_cloud_forwarding_class Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  In Policy Based Routing(forwarding) (PBR) PBR policy can select Forwarding Class object as action When match condition is satisfied(true). Forwarding class determines three things:    Ordered list ...
---

# f5_distributed_cloud_forwarding_class (Data Source)

In Policy Based Routing(forwarding) (PBR) PBR policy can select Forwarding Class object as action When match condition is satisfied(true). Forwarding class determines three things:    Ordered list ...

## Example Usage

```hcl
data "f5_distributed_cloud_forwarding_class" "example" {
  name      = "example-forwarding_class"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
