---
page_title: "f5_distributed_cloud_tunnel Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Tunnel configuration allows user to specify parameters for configuring static tunnels. Configuration involves specification of encapsulation and related parameters to be used for this tunnel Payloa...
---

# f5_distributed_cloud_tunnel (Resource)

Tunnel configuration allows user to specify parameters for configuring static tunnels. Configuration involves specification of encapsulation and related parameters to be used for this tunnel Payloa...

## Example Usage

```hcl
resource "f5_distributed_cloud_tunnel" "example" {
  name        = "example-tunnel"
  namespace   = "system"
  description = "Example Tunnel resource"
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

Tunnel can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_tunnel.example namespace/name
```
