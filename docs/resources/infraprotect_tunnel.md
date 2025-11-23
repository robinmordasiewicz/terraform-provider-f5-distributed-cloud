---
page_title: "f5_distributed_cloud_infraprotect_tunnel Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit tunnel information
---

# f5_distributed_cloud_infraprotect_tunnel (Resource)

DDoS transit tunnel information

## Example Usage

```hcl
resource "f5_distributed_cloud_infraprotect_tunnel" "example" {
  name        = "example-infraprotect_tunnel"
  namespace   = "system"
  description = "Example InfraprotectTunnel resource"
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

InfraprotectTunnel can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_infraprotect_tunnel.example namespace/name
```
