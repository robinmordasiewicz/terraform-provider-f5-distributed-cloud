---
page_title: "f5_distributed_cloud_tunnel Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Tunnel configuration allows user to specify parameters for configuring static tunnels. Configuration involves specification of encapsulation and related parameters to be used for this tunnel Payloa...
---

# f5_distributed_cloud_tunnel (Data Source)

Tunnel configuration allows user to specify parameters for configuring static tunnels. Configuration involves specification of encapsulation and related parameters to be used for this tunnel Payloa...

## Example Usage

```hcl
data "f5_distributed_cloud_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
