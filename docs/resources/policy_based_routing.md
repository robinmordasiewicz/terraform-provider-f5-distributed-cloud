---
page_title: "f5_distributed_cloud_policy_based_routing Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Policy based routing is used to control how different classes of traffic is forwarded and QOS is applied over WAN interfaces in SDWAN scenarios.
---

# f5_distributed_cloud_policy_based_routing (Resource)

Policy based routing is used to control how different classes of traffic is forwarded and QOS is applied over WAN interfaces in SDWAN scenarios.

## Example Usage

```hcl
resource "f5_distributed_cloud_policy_based_routing" "example" {
  name        = "example-policy_based_routing"
  namespace   = "system"
  description = "Example PolicyBasedRouting resource"
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

PolicyBasedRouting can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_policy_based_routing.example namespace/name
```
