---
page_title: "f5_distributed_cloud_virtual_host Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual host is main anchor configuration for a proxy. Primary application for virtual host configuration is reverse proxy.  Virtual host object is used to create a Loadbalancer, virtual service Or...
---

# f5_distributed_cloud_virtual_host (Resource)

Virtual host is main anchor configuration for a proxy. Primary application for virtual host configuration is reverse proxy.  Virtual host object is used to create a Loadbalancer, virtual service Or...

## Example Usage

```hcl
resource "f5_distributed_cloud_virtual_host" "example" {
  name        = "example-virtual_host"
  namespace   = "system"
  description = "Example VirtualHost resource"
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

VirtualHost can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_virtual_host.example namespace/name
```
