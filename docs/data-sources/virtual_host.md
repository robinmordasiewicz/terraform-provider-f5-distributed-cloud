---
page_title: "f5xc_virtual_host Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual host is main anchor configuration for a proxy. Primary application for virtual host configuration is reverse proxy.  Virtual host object is used to create a Loadbalancer, virtual service Or...
---

# f5xc_virtual_host (Data Source)

Virtual host is main anchor configuration for a proxy. Primary application for virtual host configuration is reverse proxy.  Virtual host object is used to create a Loadbalancer, virtual service Or...

## Example Usage

```hcl
data "f5xc_virtual_host" "example" {
  name      = "example-virtual_host"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
