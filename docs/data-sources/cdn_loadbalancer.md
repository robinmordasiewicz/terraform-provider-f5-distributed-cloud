---
page_title: "f5_distributed_cloud_cdn_loadbalancer Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  CDN Loadbalancer view defines a required parameters that can be used in CRUD, to create and manage CDN loadbalancer. It can be used to create CDN loadbalancer and HTTPS loadbalancer.  View will cre...
---

# f5_distributed_cloud_cdn_loadbalancer (Data Source)

CDN Loadbalancer view defines a required parameters that can be used in CRUD, to create and manage CDN loadbalancer. It can be used to create CDN loadbalancer and HTTPS loadbalancer.  View will cre...

## Example Usage

```hcl
data "f5_distributed_cloud_cdn_loadbalancer" "example" {
  name      = "example-cdn_loadbalancer"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
