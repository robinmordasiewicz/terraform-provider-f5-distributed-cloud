---
page_title: "f5xc_cdn_loadbalancer Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  CDN Loadbalancer view defines a required parameters that can be used in CRUD, to create and manage CDN loadbalancer. It can be used to create CDN loadbalancer and HTTPS loadbalancer.  View will cre...
---

# f5xc_cdn_loadbalancer (Resource)

CDN Loadbalancer view defines a required parameters that can be used in CRUD, to create and manage CDN loadbalancer. It can be used to create CDN loadbalancer and HTTPS loadbalancer.  View will cre...

## Example Usage

```hcl
resource "f5xc_cdn_loadbalancer" "example" {
  name        = "example-cdn_loadbalancer"
  namespace   = "system"
  description = "Example CDNLoadbalancer resource"
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

CDNLoadbalancer can be imported using the namespace and name:

```shell
terraform import f5xc_cdn_loadbalancer.example namespace/name
```
