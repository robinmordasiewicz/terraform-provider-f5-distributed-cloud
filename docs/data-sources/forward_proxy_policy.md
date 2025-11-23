---
page_title: "f5_distributed_cloud_forward_proxy_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Forward Proxy policy defines access control rules for connections going via forward Proxy. It is view type of config object that uses service policy mechanism to achieve Required functionality.  Vi...
---

# f5_distributed_cloud_forward_proxy_policy (Data Source)

Forward Proxy policy defines access control rules for connections going via forward Proxy. It is view type of config object that uses service policy mechanism to achieve Required functionality.  Vi...

## Example Usage

```hcl
data "f5_distributed_cloud_forward_proxy_policy" "example" {
  name      = "example-forward_proxy_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
